package fileprocess

import (
	"fmt"
	"io"
	"lightyear/core/global"
	"lightyear/models"
	"lightyear/schemas/response"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUploadHandler(c *gin.Context) {

	workflowID, err := parseWorkflowID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}
	// if user has access to this workflow
	// ! currently, if user is the owner of workflow => user has access
	userID := c.GetUint("UserID")
	workflow, err := models.GetWorkflowByWorkflowID(workflowID)
	if err != nil {
		response.Fail(0, nil, "error: failed to get workflow", c)
	}

	// validate if user has access to this workflow
	if workflow.OwnerID != userID {
		response.Fail(0, nil, "error: user does not have access to this workflow", c)
	}

	// get file from request
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(0, nil, "file upload error", c)
		return
	}

	originalExtension := filepath.Ext(file.Filename)
	originalNameWithoutExtension := strings.TrimSuffix(file.Filename, originalExtension)

	originalImageBasePath := global.Config.Storage.OriginalImagePath

	uuid := uuid.NewString()

	/* Save File Info To DB */
	fileInDB := &models.File{
		UUID:                 uuid,
		NameWithoutExtension: originalNameWithoutExtension,
		OriginalExtension:    originalExtension,
		FileStatus:           models.FileUploaded,
		OwnerID:              userID,
		WorkflowID:           workflowID,
	}
	err = models.CreateFile(fileInDB)
	if err != nil {
		response.Fail(0, nil, "error: failed to save file info into db", c)
		return
	}

	/* Save File To Local Storage */
	err = saveMultipartFile(file, originalImageBasePath, uuid, originalExtension)

	if err != nil {
		response.Fail(0, nil, "file upload error", c)
		return
	}

	response.Ok(fileInDB.UUID, "success: file uploaded successfully", c)

}

func saveMultipartFile(file *multipart.FileHeader, basePath string, prefixName string, extensionName string) (err error) {

	fileName := prefixName + extensionName

	filePath := filepath.Join(basePath, fileName)

	multipartFile, err := file.Open()
	if err != nil {
		return err
	}
	defer multipartFile.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, multipartFile); err != nil {
		return err
	}
	return nil

}

// parses workflow_id from request's path parameter
func parseWorkflowID(c *gin.Context) (uint, error) {
	workflowID := c.Param("workflow_id")
	if workflowID == "" {
		err := fmt.Errorf("error: workflow_id not provided")
		return 0, err
	}
	workflowIDInt, err := strconv.ParseUint(workflowID, 10, 64)
	if err != nil {
		err = fmt.Errorf("error: failed to parse workflow_id")
		return 0, err
	}
	return uint(workflowIDInt), nil
}
