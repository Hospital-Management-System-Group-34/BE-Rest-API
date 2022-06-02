package jwt

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
