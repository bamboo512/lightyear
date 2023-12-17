package workflow

import (
	"fmt"
	"lightyear/models"
	"lightyear/schemas/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type WorkflowCreateRequest struct {
	Name string `json:"name"`
}

type WorkflowUpdateRequest struct {
	Name string `json:"name"`
}

func WorkflowCreateHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	var request WorkflowCreateRequest
	var workflow models.Workflow
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.Fail(0, nil, "failure: failed to parse request body", c)
		return
	}

	// if does not have a name
	if request.Name == "" {
		workflow.Name = "Workflow " + time.Now().Format("2006-01-02 15:04:05")
	}

	workflow.OwnerID = userID
	err = models.CreateWorkflow(&workflow)
	if err != nil {
		response.Fail(0, nil, "error: failed to create workflow", c)
		return
	}

	response.Ok(workflow.ID, "success: workflow created successfully", c)
}

func ListFileHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	workflowID, err := parseWorkflowID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	pageNumber, pageSize, err := parsePageNumberAndSize(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	// check if workflow exists	and belongs to the user
	workflow, err := models.GetWorkflowByWorkflowID(workflowID)
	if err != nil {
		response.Fail(0, nil, "error: failed to get workflow by id", c)
		return
	}
	if workflow.OwnerID != userID {
		response.Fail(0, nil, "failure: this workflow does not belong to you", c)
		return
	}

	files, count, total, err := models.GetFilesByWorkflowID(workflowID, pageNumber, pageSize)
	if err != nil {
		response.Fail(0, nil, "error: failed to get files by workflow id", c)
		return
	}

	response.OkWithList(files, count, total, "success: files retrieved successfully", c)

}

func ListWorkflowHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	pageNumber, pageSize, err := parsePageNumberAndSize(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	workflows, count, total, err := models.GetWorkflowsByUserID(userID, pageNumber, pageSize)
	if err != nil {
		response.Fail(0, nil, "error: failed to get workflows by user id", c)
		return
	}

	response.OkWithList(workflows, count, total, "success: workflows retrieved successfully", c)

}

func WorkflowUpdateHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	workflowID, err := parseWorkflowID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	var workflowUpdateRequest WorkflowUpdateRequest
	var workflow models.Workflow
	err = c.ShouldBindJSON(&workflowUpdateRequest)
	if err != nil {
		response.Fail(0, nil, "failure: failed to parse request body", c)
		return
	}

	// check if workflow exists	and belongs to the user
	workflow, err = models.GetWorkflowByWorkflowID(workflowID)
	if err != nil {
		response.Fail(0, nil, "error: failed to get workflow by id", c)
		return
	}
	if workflow.OwnerID != userID {
		response.Fail(0, nil, "failure: this workflow does not belong to you", c)
		return
	}

	workflow.Name = workflowUpdateRequest.Name

	err = models.UpdateWorkflow(&workflow)
	if err != nil {
		response.Fail(0, nil, "error: failed to update workflow", c)
		return
	}

	response.Ok(workflow.ID, "success: workflow updated successfully", c)

}

func WorkflowDeleteHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	workflowID, err := parseWorkflowID(c)
	if err != nil {
		response.Fail(0, nil, err.Error(), c)
		return
	}

	// check if workflow exists	and belongs to the user
	workflow, err := models.GetWorkflowByWorkflowID(workflowID)
	if err != nil {
		response.Fail(0, nil, "error: failed to get workflow by id", c)
		return
	}
	if workflow.OwnerID != userID {
		response.Fail(0, nil, "failure: this workflow does not belong to you", c)
		return
	}

	err = models.DeleteWorkflow(&workflow)
	if err != nil {
		response.Fail(0, nil, "error: failed to delete workflow", c)
		return
	}

	response.Ok(workflow.ID, "success: workflow deleted successfully", c)

}

func parsePageNumberAndSize(c *gin.Context) (int, int, error) {
	pageNumber := c.Query("page_number")
	pageSize := c.Query("page_size")

	if pageNumber == "" {
		pageNumber = "0"
	}
	if pageSize == "" {
		pageSize = "20"
	}
	pageNumberInt, err := strconv.ParseUint(pageNumber, 10, 32)
	if err != nil {
		err = fmt.Errorf("error: failed to parse page number")
		return 0, 0, err
	}
	pageSizeInt, err := strconv.ParseUint(pageSize, 10, 32)
	if err != nil {
		err = fmt.Errorf("error: failed to parse page size")
		return 0, 0, err
	}
	return int(pageNumberInt), int(pageSizeInt), nil
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
