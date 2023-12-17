package auth

import (
	"lightyear/models"
	"lightyear/pkg/auth/utils"
	"lightyear/schemas"
	"lightyear/schemas/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LogIn(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	existingUser, err := models.GetUserByEmail(user.Email)
	// TODO: db error
	if err != nil {
		response.Fail(0, nil, "error: db ...", c)
	}

	if existingUser.ID == 0 {
		response.Fail(0, nil, "failure: user does not exist", c)
		return
	}

	errHash := utils.VerifyPassword(user.Password, existingUser.Password)

	if !errHash {
		response.Fail(0, nil, "failure: invalid password", c)
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &schemas.Claims{
		ID: existingUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   existingUser.Email,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(utils.JwtKey)

	if err != nil {
		response.Fail(0, nil, "error: could not generate token", c)
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	response.Ok(tokenString, "success: user successfully logged in", c)
}

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	var existingUser models.User

	models.GetUserByEmail(user.Email)

	if existingUser.ID != 0 {
		response.Fail(0, nil, "failure: user already exists", c)
		return
	}

	var errHash error
	user.Password, errHash = utils.HashPassword(user.Password)

	if errHash != nil {
		response.Fail(0, nil, "error: could not generate password hash", c)
		return
	}

	models.CreateUser(user.Email, user.Password)

	response.Ok(nil, "success: user successfully created", c)
}

func LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	response.Ok(nil, "success: user successfully logged out", c)
}
