package fileprocess

import (
	"fmt"
	"lightyear/core/global"
	"lightyear/models"
	"lightyear/schemas/response"
	"mime"
	"os"

	"github.com/gin-gonic/gin"
)

// TODO: Original File Download Handler
func OriginalFileDownloadHandler(c *gin.Context) {

}

// TODO: thumbnail
func ThumbnailFileDownloadHandler(c *gin.Context) {

}

func TranscodedFileDownloadHandler(c *gin.Context) {
	uuid, err := parseUUID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	file, err := models.GetFileByUUID(uuid)
	if err != nil {
		response.Fail(0, nil, "failure: file info does not exist in database", c)
		return
	}
	// if file.ID == 0 {
	// 	response.Fail(0, nil, "error: file does not exist", c)
	// }

	// TODO: can be deleted after having implemented
	// When upload canceled or failed => Remove file record from database
	if file.FileStatus < 0 {
		response.Fail(0, nil, "failure: file is not uploaded yet", c)
		return
	}

	if file.FileStatus != models.FileConverted {
		response.Fail(0, nil, "error: image has not been transcoded yet.", c)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.NameWithoutExtension+file.ExpectedExtension))

	fileBytes, err := getFileBytes("transcoded", uuid+file.ExpectedExtension)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
	}

	mimeType := mime.TypeByExtension(file.ExpectedExtension)

	c.Data(200, mimeType, fileBytes)
}

// Get UUID from download link
func parseUUID(c *gin.Context) (string, error) {

	uuid := c.Param("uuid")
	var err error
	// if is not a valid uuid
	if len(uuid) != 36 {
		err = fmt.Errorf("error: invalid file uuid")
	}
	return uuid, err

}

func getFileBytes(downloadType string, fileName string) ([]byte, error) {
	var fullFilePath string
	switch downloadType {
	case "original":
		{
			fullFilePath = global.Config.Storage.OriginalImagePath + fileName
		}
	case "thumbnail":
		{
			fullFilePath = global.Config.Storage.ThumbnailImagePath + fileName
		}
	case "transcoded":
		{
			fullFilePath = global.Config.Storage.TranscodedImagePath + fileName
		}
	}

	// Return the encoded image
	fileBytes, err := os.ReadFile(fullFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)

	}
	return fileBytes, nil
}
