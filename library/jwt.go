package library

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SecretKey string
}

type customClaims struct {
	ID   string
	Name string
	jwt.StandardClaims
}

func NewJWT(secretKey string) JWT {
	return JWT{
		SecretKey: secretKey,
	}
}

func (c JWT) CreateToken(id string, name string) (string, error) {

	expiredTime := time.Now().Add(24 * time.Hour)
	claims := &customClaims{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString([]byte(c.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (c JWT) Authentication(secretKey string) fiber.Handler {
	return func(context *fiber.Ctx) error {
		errHandle := func(err error, message string) error {
			fmt.Println(err.Error())
			return context.Status(http.StatusUnauthorized).Send([]byte(message))
		}

		input := context.Request().Header.Peek("Authorization")

		if !strings.Contains(string(input), "Bearer") {
			return errHandle(errors.New("invalid header"), "invalid header")
		}

		authToken := ""
		arrayToken := strings.Split(string(input), " ")
		if len(arrayToken) == 2 {
			authToken = arrayToken[1]
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			return errHandle(err, "token expired")
		}

		if !token.Valid {
			return errHandle(errors.New("invalid token"), "invalid token")
		}

		return context.Next()
	}
}
