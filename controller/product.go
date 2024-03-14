package controller

import (
	"encoding/json"
	"errors"
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

func (c Product) List(ctx *fiber.Ctx) (int, string, interface{}, error) {

	result, err := c.repo.PRODUCT.List()
	if err != nil {
		return http.StatusInternalServerError, "list product error", nil, err
	}

	return http.StatusOK, "succes", result, err
}

func (c Product) Get(ctx *fiber.Ctx) (int, string, interface{}, error) {

	id := ctx.Params("id")
	result, err := c.repo.PRODUCT.GetByID(id)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	return http.StatusOK, "succes", result, err
}
