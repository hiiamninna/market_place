package controller

import (
	"encoding/json"
	"fmt"
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

	// set validation here
	err = library.Validate(input)
	if err != nil {
		return http.StatusBadRequest, err.Error(), nil, err
	}

	generated, err := bcrypt.GenerateFromPassword([]byte(input.Password), c.bcryptSalt)
	if err != nil {
		return http.StatusInternalServerError, "failed generate", nil, err
	}

	input.Password = string(generated)
	input.ID = generateUUID()

	_, err = c.repo.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "User registered failed", nil, err
	}

	token, err := c.jwt.CreateToken(input.ID, input.Name)
	if err != nil {
		fmt.Println(err)
	}

	resp := collections.UserRegister{
		Name:        input.Name,
		Username:    input.Username,
		AccessToken: token,
	}

	return http.StatusCreated, "User registered successfully", resp, err
}
