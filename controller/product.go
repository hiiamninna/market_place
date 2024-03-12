package controller

import (
	"encoding/json"
	"market_place/collections"
	"market_place/repository"
	"net/http"

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

	input.ID = generateUUID()

	_, err = c.repo.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "product added failed", nil, err
	}

	return http.StatusOK, "product added successfully", nil, err
}
