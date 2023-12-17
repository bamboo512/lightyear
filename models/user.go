package models

import (
	"lightyear/core/global"
	"time"

	"gorm.io/gorm"
)

type Role int

const (
	FreeUser Role = iota
	ProUser  Role = iota
	Admin    Role = iota
)

type User struct {
	gorm.Model
	Username          string     `gorm:"type:varchar(16);unique" json:"username"`
	Email             string     `gorm:"type:varchar(256)" json:"email"`
	Password          string     `gorm:"type:varchar(64)" json:"password"`
	Workflows         []Workflow `json:"workflows" gorm:"foreignKey:OwnerID"` // 1:N relationship
	Role              Role       `gorm:"type:integer" json:"role"`
	SubscriptionUntil time.Time  `gorm:"type:date" json:"subscription_until"`
}

func CreateUser(email string, password string) error {
	user := User{
		Email:             email,
		Password:          password,
		SubscriptionUntil: time.Now().AddDate(0, 1, 0),
	}
	err := global.DB.Create(&user).Error
	return err
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := global.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
func GetUserByEmail(email string) (User, error) {
	var user User
	err := global.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserByUserID(id uint) (User, error) {
	var user User
	err := global.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// Delete User
func DeleteUserByID(id uint) error {
	err := global.DB.Where("id = ?", id).Delete(&User{}).Error
	return err
}

// Change User Info
func ChangeUserInfoByID(id uint, username string, password string) error {
	err := global.DB.Model(&User{}).Where("id = ?", id).Updates(User{Username: username, Password: password}).Error
	return err
}
