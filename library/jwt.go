package library

import "github.com/golang-jwt/jwt"

type JWT struct {
	SecretKey string
}

type customClaims struct {
	ID   string
	Name string
	jwt.StandardClaims
}
