package controller

import (
	"market_place/library"
	"market_place/repository"

	"github.com/google/uuid"
)

type Controller struct {
	USER    User
	PRODUCT Product
	IMAGE   Image
}

func NewController(repo repository.Repository, jwt library.JWT, bcryptSalt int, s3 library.S3) Controller {
	return Controller{
		USER:    NewUserController(repo.USER, jwt, bcryptSalt),
		PRODUCT: NewProductController(repo.PRODUCT),
		IMAGE:   NewImageController(s3),
	}
}

func generateUUID() string {
	return uuid.NewString()
}
