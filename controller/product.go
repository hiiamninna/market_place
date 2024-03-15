package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"market_place/collections"
	"market_place/library"
	"market_place/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	repo repository.Repository
}

func NewProductController(repo repository.Repository) Product {
	return Product{
		repo: repo,
	}
}

func (c Product) Create(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.ProductInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	input.UserID, _ = library.GetUserID(ctx)

	err = library.Validate(input)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil, err
	}

	err = c.repo.PRODUCT.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "product added failed", nil, err
	}

	return http.StatusOK, "product added successfully", nil, err
}

func (c Product) Update(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)

	id := ctx.Params("id")
	_, err := c.repo.PRODUCT.GetOwnByID(id, userID)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	raw := ctx.Request().Body()
	input := collections.ProductInput{}
	err = json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}
	input.ID = id

	err = library.Validate(input)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil, err
	}

	err = c.repo.PRODUCT.Update(input)
	if err != nil {
		return http.StatusInternalServerError, "product updated failed", nil, err
	}

	return http.StatusOK, "product updated successfully", nil, err
}

func (c Product) Delete(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)

	id := ctx.Params("id")
	_, err := c.repo.PRODUCT.GetOwnByID(id, userID)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	err = c.repo.PRODUCT.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, "product delete failed", nil, err
	}

	return http.StatusOK, "product deleted successfully", nil, err
}

func (c Product) UpdateStock(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.ProductStockInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	userID, _ := library.GetUserID(ctx)

	input.ID = ctx.Params("id")
	_, err = c.repo.PRODUCT.GetOwnByID(input.ID, userID)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	err = c.repo.PRODUCT.UpdateStock(input.ID, input.Stock)
	if err != nil {
		return http.StatusInternalServerError, "product update stock failed", nil, err
	}

	return http.StatusOK, "product stock update successfully", nil, err
}

func (c Product) List(ctx *fiber.Ctx) (int, string, collections.Meta, interface{}, error) {

	input := collections.ProductPageInput{}
	if err := ctx.QueryParser(&input); err != nil {
		return http.StatusInternalServerError, "list product error", collections.Meta{}, nil, err
	}

	input.UserID, _ = library.GetUserID(ctx)

	result, err := c.repo.PRODUCT.List(input)
	if err != nil {
		return http.StatusInternalServerError, "list product error", collections.Meta{}, nil, err
	}

	totalRow, err := c.repo.PRODUCT.CountList(input)
	if err != nil {
		return http.StatusInternalServerError, "list product error", collections.Meta{}, nil, err
	}

	meta := collections.Meta{
		Limit:  input.Limit,
		Offset: input.Offset,
		Total:  totalRow,
	}

	return http.StatusOK, "ok", meta, result, err
}

func (c Product) Get(ctx *fiber.Ctx) (int, string, interface{}, error) {

	var err error
	productDetail := collections.ProductDetail{}

	id := ctx.Params("id")
	productDetail.Product, err = c.repo.PRODUCT.GetByID(id)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, fmt.Errorf("get by id : %w", err)
	}

	productDetail.Product.PurchaseCount, err = c.repo.PAYMENT.GetPurchaseCountByProductID(id)
	if err != nil {
		fmt.Errorf("get purchase count : %w", err)
	}

	seller, err := c.repo.USER.GetByID(productDetail.Product.UserID)
	if err != nil {
		fmt.Errorf("get seller : %w", err)
	}

	productDetail.Seller.Name = seller.Name
	productDetail.Seller.ProductSoldTotal, err = c.repo.PAYMENT.GetProductSoldTotalByUser(seller.ID)
	if err != nil {
		fmt.Errorf("get product sold total : %w", err)
	}

	productDetail.Seller.BankAccounts, err = c.repo.BANK_ACCOUNT.List(seller.ID)
	if err != nil {
		fmt.Errorf("get bank accounts : %w", err)
	}

	return http.StatusOK, "ok", productDetail, err
}
