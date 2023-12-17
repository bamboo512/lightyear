package workflow

import (
	"lightyear/pkg/auth/middleware"

	"github.com/gin-gonic/gin"
)

func AddRoutersTo(router *gin.Engine) {

	group := router.Group("/workflow")
	group.Use(middleware.RequiresAuthorized())

	group.PUT("/create", WorkflowCreateHandler)
	group.GET("/get_file_list/:workflow_id", ListFileHandler)
	group.GET("/get_workflow_list", ListWorkflowHandler)
	group.POST("/update/:workflow_id", WorkflowUpdateHandler)
	group.DELETE("/delete/:workflow_id", WorkflowDeleteHandler)
}
