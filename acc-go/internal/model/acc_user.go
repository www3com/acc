package model

import (
	"github.com/upbos/go-saber/db"
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

type UserDao struct{}

func (d *UserDao) Insert(tx *gorm.DB, user *User) error {
	return tx.Create(user).Error
}

func (d *UserDao) Get(username string) (user *User, err error) {
	err = db.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
