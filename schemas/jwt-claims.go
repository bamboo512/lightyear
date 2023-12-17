package schemas

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}
