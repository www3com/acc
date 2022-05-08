package model

import (
	"accounting-service/pkg/db"
	"gorm.io/gorm"
)

type User struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	CreateTime int64  `json:"-"`
	UpdateTime int64  `json:"-"`
	Nickname   string `json:"nickname"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Agree      int    `json:"agree"`
}

func (User) TableName() string {
	return "upm_user"
}

func InsertUser(tx *gorm.DB, user *User) error {
	if err := tx.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
