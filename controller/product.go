package controller

import (
	"encoding/json"
	"errors"
	"market_place/collections"
	"market_place/library"
	"market_place/repository"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	repo repository.Product
}

func NewProductController(product repository.Product) Product {
	return Product{
		repo: product,
	}
}

func (c Product) Create(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.ProductInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	maps, err := library.GetAllSession(ctx)
	if err != nil {
		return http.StatusBadRequest, "failed get session", nil, err
	}

	input.UserID, _ = strconv.Atoi(maps[`user_id`].(string))

	err = c.repo.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "product added failed", nil, err
	}

	return http.StatusOK, "product added successfully", nil, err
}

func (c Product) Update(ctx *fiber.Ctx) (int, string, interface{}, error) {

	id := ctx.Params("id")
	_, err := c.repo.GetByID(id)
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

	err = c.repo.Update(input)
	if err != nil {
		return http.StatusInternalServerError, "product updated failed", nil, err
	}

	return http.StatusOK, "product updated successfully", nil, err
}

func (c Product) Delete(ctx *fiber.Ctx) (int, string, interface{}, error) {

	id := ctx.Params("id")
	_, err := c.repo.GetByID(id)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	err = c.repo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, "product delete failed", nil, err
	}

	return http.StatusOK, "product deleted successfully", nil, err
}

func (c Product) List(ctx *fiber.Ctx) (int, string, interface{}, error) {

	result, err := c.repo.List()
	if err != nil {
		return http.StatusInternalServerError, "list product error", nil, err
	}

	return http.StatusOK, "succes", result, err
}

func (c Product) Get(ctx *fiber.Ctx) (int, string, interface{}, error) {

	id := ctx.Params("id")
	result, err := c.repo.GetByID(id)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	return http.StatusOK, "succes", result, err
}

func (c Product) UpdateStock(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.ProductStockInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	input.ID = ctx.Params("id")
	_, err = c.repo.GetByID(input.ID)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, errors.New("product not found")
	}

	err = c.repo.UpdateStock(input.ID, input.Stock)
	if err != nil {
		return http.StatusInternalServerError, "product update stock failed", nil, err
	}

	return http.StatusOK, "product stock update successfully", nil, err
}
