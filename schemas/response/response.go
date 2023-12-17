package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ListResponse[T any] struct {
	List  []T   `json:"list"`
	Count int64 `json:"count"`
	Total int64 `json:"total"`
}

const (
	ERROR   = 0
	SUCCESS = 1
)

// Return JSON response
func Ok(data any, message string, c *gin.Context) {
	respondWithJSON(SUCCESS, data, message, c)
}

// Return JSON response with list
func OkWithList[T any](list []T, count int64, total int64, message string, c *gin.Context) {
	respondWithJSON(
		SUCCESS,
		ListResponse[T]{
			list,
			count,
			total,
		},
		message,
		c,
	)
}

func Fail(code int, data any, message string, c *gin.Context) {
	respondWithJSON(code, data, message, c)
}

func respondWithJSON(code int, data any, message string, c *gin.Context) {

	c.JSON(200,
		Response{
			code,
			data,
			message,
		})

}
