package database

import (
	"lightyear/core/global"
	"lightyear/models"
	"os"
)

func MigrateDatabase() error {

	err := global.DB.AutoMigrate(&models.User{}, &models.Workflow{}, &models.File{})
	if err != nil {
		global.Logger.Error("migrate database failed: ", err)
		os.Exit(1)
	}
	return err
}
