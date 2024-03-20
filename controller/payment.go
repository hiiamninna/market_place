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

type Payment struct {
	repo repository.Repository
}

func NewPaymentController(repo repository.Repository) Payment {
	return Payment{
		repo: repo,
	}
}

func (c Payment) Create(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.PaymentInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	message, err := library.ValidateInput(input)
	if err != nil {
		return http.StatusBadRequest, message, nil, err
	}

	input.UserID, _ = library.GetUserID(ctx)
	if input.UserID == "" {
		return http.StatusForbidden, "please check your credential", nil, errors.New("not login")
	}

	_, err = c.repo.BANK_ACCOUNT.GetByID(input.BankAccountID, input.UserID)
	if err != nil {
		return http.StatusBadRequest, "incorrect payment information", nil, err
	}

	input.ProductID = ctx.Params("id")
	product, err := c.repo.PRODUCT.GetByID(input.ProductID)
	if err != nil {
		return http.StatusNotFound, "product not found", nil, err
	}

	if input.Quantity > product.Stock {
		return http.StatusBadRequest, "insufficient quantity", nil, err
	}

	input.TotalPayment = product.Price * input.Quantity

	err = c.repo.PAYMENT.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "payment processed failed", nil, err
	}

	return http.StatusOK, "payment processed successfully", nil, err
}
