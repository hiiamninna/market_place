package controller

import (
	"market_place/library"
	"market_place/repository"

	"github.com/google/uuid"
)

type Controller struct {
	USER User
}

func NewController(repo repository.Repository, jwt library.JWT, bcryptSalt int) Controller {
	return Controller{
		USER: NewUserController(repo.USER, jwt, bcryptSalt),
	}
}

func generateUUID() string {
	return uuid.NewString()
}
