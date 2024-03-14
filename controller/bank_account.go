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

type BankAccount struct {
	repo repository.Repository
}

func NewBankAccountRepository(repo repository.Repository) BankAccount {
	return BankAccount{
		repo: repo,
	}
}

func (c BankAccount) Create(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.BankAccountInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	input.UserID, _ = library.GetUserID(ctx)

	err = c.repo.BANK_ACCOUNT.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "bank account added failed", nil, err
	}

	return http.StatusOK, "bank account added successfully", nil, err
}

func (c BankAccount) Update(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)

	id := ctx.Params("id")
	_, err := c.repo.BANK_ACCOUNT.GetByID(id, userID)
	if err != nil {
		return http.StatusNotFound, "bank account not found", nil, errors.New("bank account not found")
	}

	raw := ctx.Request().Body()
	input := collections.BankAccountInput{}
	err = json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}
	input.ID = id

	err = c.repo.BANK_ACCOUNT.Update(input)
	if err != nil {
		return http.StatusInternalServerError, "bank account updated failed", nil, err
	}

	return http.StatusOK, "bank account updated successfully", nil, err
}

func (c BankAccount) Delete(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)

	id := ctx.Params("id")
	_, err := c.repo.BANK_ACCOUNT.GetByID(id, userID)
	if err != nil {
		return http.StatusNotFound, "bank account not found", nil, errors.New("bank account not found")
	}

	err = c.repo.BANK_ACCOUNT.Delete(id, userID)
	if err != nil {
		return http.StatusInternalServerError, "bank account delete failed", nil, err
	}

	return http.StatusOK, "bank account deleted successfully", nil, err
}

func (c BankAccount) List(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)

	result, err := c.repo.BANK_ACCOUNT.List(userID)
	if err != nil {
		return http.StatusInternalServerError, "list bank account error", nil, err
	}

	return http.StatusOK, "success", result, err
}
