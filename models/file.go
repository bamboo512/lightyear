package models

import (
	"lightyear/core/global"

	"gorm.io/gorm"
)

type Status int

const (
	FileError      Status = -2
	FileCanceled   Status = -1
	FileUploaded   Status = 0
	FileConverted  Status = 1
	FileDownloaded Status = 2
)

type File struct {
	gorm.Model
	UUID                 string `gorm:"type:varchar(36);unique" json:"uuid"`
	NameWithoutExtension string `gorm:"type:varchar(256)" json:"name_without_extension"` // 假设原始名称可以更长，并且不是唯一的NameWithoutExtension
	OriginalExtension    string `gorm:"type:varchar(12)" json:"original_extension"`
	ExpectedExtension    string `gorm:"type:varchar(12)" json:"expected_extension"`
	FileStatus           Status `gorm:"type:integer" json:"status"` // 状态是一个整数
	WorkflowID           uint   `gorm:"type:integer" json:"workflow_id"`
	OwnerID              uint   `gorm:"type:integer" json:"owner_id"`
}

func GetFileByUUID(uuid string) (File, error) {
	var file File
	err := global.DB.Where("uuid = ?", uuid).First(&file).Error

	return file, err
}

func CreateFile(file *File) error {
	err := global.DB.Create(&file).Error
	return err
}

func UpdateFile(file *File) error {
	err := global.DB.Save(&file).Error
	return err
}

func DeleteFile(file *File) error {
	err := global.DB.Delete(&file).Error
	return err
}

func GetFilesByWorkflowID(workflowID uint, pageNumber int, pageSize int) ([]File, int64, int64, error) {

	var files []File

	var total int64 = 0

	// 使用 model 参数来提供模型的类型信息。
	err := global.DB.Model(&File{}).
		Where("workflow_id = ?", workflowID).
		Count(&total).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		Find(&files).
		Error

	if err != nil {
		return nil, 0, 0, err
	}

	count := int64(len(files))

	return files, count, total, err
}
