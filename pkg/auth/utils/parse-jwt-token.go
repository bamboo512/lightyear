// PATH: go-auth/utils/ParseToken.go

package utils

import (
	"lightyear/core/global"
	"lightyear/schemas"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret_key_should_be_complicated")

func ParseJwtToken(jwtString string) (claims *schemas.Claims, err error) {
	token, err := jwt.ParseWithClaims(jwtString, &schemas.Claims{}, func(token *jwt.Token) (any, error) {
		return JwtKey, nil
	})

	if err != nil {
		global.Logger.Debugln(err)
		return nil, err
	}

	claims, ok := token.Claims.(*schemas.Claims)

	if !ok {

		return nil, err
	}

	return claims, nil
}
