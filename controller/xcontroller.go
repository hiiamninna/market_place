package controller

import (
	"github.com/hiiamninna/market_place/library"
	"github.com/hiiamninna/market_place/repository"

	"github.com/google/uuid"
)

type Controller struct {
	USER         User
	PRODUCT      Product
	IMAGE        Image
	BANK_ACCOUNT BankAccount
	PAYMENT      Payment
}

func NewController(repo repository.Repository, jwt library.JWT, bcryptSalt int, s3 library.S3) Controller {
	return Controller{
		USER:         NewUserController(repo, jwt, bcryptSalt),
		PRODUCT:      NewProductController(repo),
		IMAGE:        NewImageController(s3),
		BANK_ACCOUNT: NewBankAccountController(repo),
		PAYMENT:      NewPaymentController(repo),
	}
}

func generateUUID() string {
	return uuid.NewString()
}
