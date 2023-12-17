package auth

import "github.com/gin-gonic/gin"

func AddRoutersTo(c *gin.Engine) {
	auth := c.Group("/auth")

	auth.POST("/login", LogIn)
	auth.POST("/signup", SignUp)
	auth.POST("/logout", LogOut)
	// auth.POST("/refresh", Refresh)

}
