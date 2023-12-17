package middleware

import (
	"lightyear/core/global"
	"lightyear/pkg/auth/utils"
	"lightyear/schemas/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequiresAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			global.Logger.Errorln("error: token is empty")
			response.Fail(0, nil, "unarthorinzed: token is empty", c)
			c.Abort()
			return
		}

		token := strings.Split(bearerToken, " ")[1]
		global.Logger.Warnln(token)
		claims, err := utils.ParseJwtToken(token)

		if err != nil {
			response.Fail(0, nil, "unauthorized: failed to parse jwt token, "+err.Error(), c)
			c.Abort()
			return
		}

		c.Set("UserID", claims.ID)
		c.Next()
	}
}

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Request.Header.Get("Authorization")

		if cookie == "" {
			response.Fail(0, nil, "unauthorized", c)
			c.Abort()
			return
		}

		claims, err := utils.ParseJwtToken(cookie)

		if err != nil {
			response.Fail(0, nil, "unauthorized", c)
			c.Abort()
			return
		}

		c.Set("ID", claims.ID)
		c.Next()
	}
}
