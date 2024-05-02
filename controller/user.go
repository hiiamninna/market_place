package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hiiamninna/market_place/collections"
	"github.com/hiiamninna/market_place/library"
	"github.com/hiiamninna/market_place/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo       repository.Repository
	jwt        library.JWT
	bcryptSalt int
}

func NewUserController(repo repository.Repository, jwt library.JWT, bcryptSalt int) User {
	return User{
		repo:       repo,
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

	message, err := library.ValidateInput(input)
	if err != nil {
		return http.StatusBadRequest, message, nil, err
	}

	existUser, _ := c.repo.USER.GetByUsername(input.Username)
	if existUser.ID != "" {
		return http.StatusConflict, "username already exist", nil, errors.New("username existed")
	}

	generated, err := bcrypt.GenerateFromPassword([]byte(input.Password), c.bcryptSalt)
	if err != nil {
		return http.StatusInternalServerError, "failed generate", nil, err
	}

	input.Password = string(generated)

	id, err := c.repo.USER.Create(input)
	if err != nil {
		return http.StatusInternalServerError, "User registered failed", nil, err
	}

	token, err := c.jwt.CreateToken(strconv.Itoa(id), input.Name)
	if err != nil {
		return http.StatusInternalServerError, "User registered failed", nil, err
	}

	resp := collections.UserRegisterAndLogin{
		Name:        input.Name,
		Username:    input.Username,
		AccessToken: token,
	}

	return http.StatusCreated, "User registered successfully", resp, err
}

func (c User) Login(ctx *fiber.Ctx) (int, string, interface{}, error) {

	raw := ctx.Request().Body()

	input := collections.UserLoginInput{}
	err := json.Unmarshal([]byte(raw), &input)
	if err != nil {
		return http.StatusBadRequest, "unmarshal input", nil, err
	}

	message, err := library.ValidateInput(input)
	if err != nil {
		return http.StatusBadRequest, message, nil, err
	}

	user, err := c.repo.USER.GetByUsername(input.Username)
	if err != nil {
		return http.StatusNotFound, "User not found", nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return http.StatusBadRequest, "Check your password again", nil, err
	}

	token, err := c.jwt.CreateToken(user.ID, user.Name)
	if err != nil {
		fmt.Println(err)
	}

	resp := collections.UserRegisterAndLogin{
		Name:        user.Name,
		Username:    user.Username,
		AccessToken: token,
	}

	return http.StatusOK, "User logged successfully", resp, err
}
