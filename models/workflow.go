package models

import (
	"lightyear/core/global"

	"gorm.io/gorm"
)

type Workflow struct {
	gorm.Model
	OwnerID uint   `gorm:"type:integer" json:"owner_id"`
	Name    string `gorm:"type:varchar(256)" json:"name"`
	Files   []File `json:"files"` // 1:N relationship
}

// create workflow
func CreateWorkflow(workflow *Workflow) error {
	err := global.DB.Create(&workflow).Error
	return err
}

// get workflow by workflowID
func GetWorkflowByWorkflowID(workflowID uint) (Workflow, error) {
	var workflow Workflow
	err := global.DB.Where("id = ?", workflowID).First(&workflow).Error
	return workflow, err
}

// get workflow list by userID
func GetWorkflowsByUserID(userID uint, pageNumber int, pageSize int) ([]Workflow, int64, int64, error) {

	var workflows []Workflow

	var total int64 = 0

	// 使用 model 参数来提供模型的类型信息。
	err := global.DB.Model(&Workflow{}).
		Where("owner_id = ?", userID).
		Count(&total).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Find(&workflows).
		Error

	if err != nil {
		return nil, 0, 0, err
	}

	count := int64(len(workflows))

	return workflows, count, total, err
}

// update workflow
func UpdateWorkflow(workflow *Workflow) error {
	err := global.DB.Save(&workflow).Error
	return err
}

// delete workflow
func DeleteWorkflow(workflow *Workflow) error {
	err := global.DB.Delete(&workflow).Error
	return err
}
