package controller

import (
	"market_place/library"
	"market_place/repository"
)

type Controller struct {
	USER User
}

func NewController(repo repository.Repository, jwt library.JWT, bcryptSalt int) Controller {
	return Controller{
		USER: NewUserController(repo.USER, jwt, bcryptSalt),
	}
}
