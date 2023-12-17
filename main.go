package main

import (
	"lightyear/cmd"
	"lightyear/pkg/auth"
	fileprocess "lightyear/pkg/file-process"
	"lightyear/pkg/workflow"

	"github.com/gin-gonic/gin"
)

const (
	DefaultQuality       = 80
	ParamErrorQuality    = "parameter error: quality can only be between 1-100"
	ParamErrorEncoding   = "parameter error: encoding can only be avif, heic or web"
	InternalServerErr    = "Internal Server Error"
	SupportedEncodings   = "avif, heic, webp"
	OriginalImagePathEnv = "ORIGINAL_IMAGE_PATH"
	EncodedImagePathEnv  = "ENCODED_IMAGE_PATH"
)

func main() {
	cmd.InitialApp()

	router := gin.Default()

	auth.AddRoutersTo(router)
	fileprocess.AddRoutersTo(router)
	workflow.AddRoutersTo(router)

	router.Run(":8080")

}
