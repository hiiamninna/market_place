package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hiiamninna/market_place/collections"
	"github.com/hiiamninna/market_place/library"
	"github.com/hiiamninna/market_place/repository"

	"github.com/gofiber/fiber/v2"
)

type BankAccount struct {
	repo repository.Repository
}

func NewBankAccountController(repo repository.Repository) BankAccount {
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

	message, err := library.ValidateInput(input)
	if err != nil {
		return http.StatusBadRequest, message, nil, err
	}

	input.UserID, _ = library.GetUserID(ctx)
	if input.UserID == "" {
		return http.StatusForbidden, "please check your credential", nil, errors.New("not login")
	}

	err = c.repo.BANK_ACCOUNT.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "bank account added failed", nil, err
	}

	return http.StatusOK, "bank account added successfully", nil, err
}

func (c BankAccount) Update(ctx *fiber.Ctx) (int, string, interface{}, error) {

	userID, _ := library.GetUserID(ctx)
	if userID == "" {
		return http.StatusForbidden, "please check your credential", nil, errors.New("not login")
	}

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

	message, err := library.ValidateInput(input)
	if err != nil {
		return http.StatusBadRequest, message, nil, err
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
	if userID == "" {
		return http.StatusForbidden, "please check your credential", nil, errors.New("not login")
	}

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
	if userID == "" {
		return http.StatusForbidden, "please check your credential", nil, errors.New("not login")
	}

	result, err := c.repo.BANK_ACCOUNT.List(userID)
	if err != nil {
		return http.StatusInternalServerError, "list bank account error", nil, err
	}

	return http.StatusOK, "success", result, err
}
