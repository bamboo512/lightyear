package fileprocess

import (
	"fmt"
	"lightyear/pkg/tasks"
	"lightyear/schemas/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ImageTranscodeHandler(c *gin.Context) {
	quality, err := getQuality(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamErrorQuality)
		return
	}

	encoding, err := getEncoding(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ParamErrorEncoding)
		return
	}

	uuid, err := getFileUUID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
	}

	tasks.AddNewTask(uuid, encoding, quality)

	response.Ok(nil, "success: image successfully transcoded", c)

}

func getQuality(c *gin.Context) (int, error) {
	qualityStr := c.DefaultQuery("quality", strconv.Itoa(DefaultQuality))
	quality, err := strconv.Atoi(qualityStr)
	if quality <= 0 || quality > 100 {
		err = fmt.Errorf(ParamErrorQuality)
	}
	return quality, err
}

func getEncoding(c *gin.Context) (string, error) {
	encoding := c.Query("encoding")
	var err error
	if encoding != "avif" && encoding != "heic" && encoding != "webp" {
		err = fmt.Errorf(ParamErrorEncoding)
	}
	return encoding, err
}

func getFileUUID(c *gin.Context) (uuid string, err error) {
	uuid = strings.Trim(c.Param("uuid"), "/")
	if uuid == "" {
		err = fmt.Errorf("error: empty file uuid")
		return "", err
	}

	return uuid, nil
}
