package model

import (
	"acc/internal/pkg/db"
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
	Agreement  int    `json:"agreement"`
	State      int
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
	if err == nil {
		return &user, nil
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return nil, err
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
