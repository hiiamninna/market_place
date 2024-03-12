package controller

import (
	"market_place/library"
	"market_place/repository"

	"github.com/google/uuid"
)

type Controller struct {
	USER    User
	PRODUCT Product
}

func NewController(repo repository.Repository, jwt library.JWT, bcryptSalt int) Controller {
	return Controller{
		USER:    NewUserController(repo.USER, jwt, bcryptSalt),
		PRODUCT: NewProductController(repo.PRODUCT),
	}
}

func generateUUID() string {
	return uuid.NewString()
}
