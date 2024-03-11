package controller

import (
	"encoding/json"
	"market_place/collections"
	"market_place/library"
	"market_place/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo       repository.User
	jwt        library.JWT
	bcryptSalt int
}

func NewUserController(user repository.User, jwt library.JWT, bcryptSalt int) User {
	return User{
		repo:       user,
		jwt:        jwt,
		bcryptSalt: bcryptSalt,
	}
}

func (c User) Register(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.InputUserRegister{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	err = library.Validate(input)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil, err
	}

	// set validation here
	generated, err := bcrypt.GenerateFromPassword([]byte(input.Password), c.bcryptSalt)
	if err != nil {
		return http.StatusInternalServerError, "failed generate", nil, err
	}

	code, message, resp, err := c.repo.Register(input.Name, input.Username, string(generated))

	return code, message, resp, err
}
