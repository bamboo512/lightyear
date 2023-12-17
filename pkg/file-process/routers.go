package fileprocess

import (
	"lightyear/pkg/auth/middleware"

	"github.com/gin-gonic/gin"
)

func AddRoutersTo(r *gin.Engine) {
	group := r.Group("/image")
	group.Use(middleware.RequiresAuthorized())

	group.POST("/upload/:workflow_id", FileUploadHandler)

	group.GET("/download/original/:uuid", OriginalFileDownloadHandler)
	group.GET("/download/thumbnail/:uuid", ThumbnailFileDownloadHandler)
	group.GET("/download/transcoded/:uuid", TranscodedFileDownloadHandler)

	group.POST("/transcode/:uuid", ImageTranscodeHandler)

}
