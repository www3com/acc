package service

import (
	"accounting-service/internal/model"
	"accounting-service/pkg/db"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Username string
	Email    string
	Password string
	Captcha  int
}

func (u *UserService) SignInByUsername() (int64, error) {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return 0, err
	}
	if user.Password == u.Password {
		return user.ID, nil
	} else {
		return 0, nil
	}
}

func (u *UserService) SignInByEmail() (int64, error) {
	return 0, nil
}

func (u *UserService) SignUp(user *model.User) (int64, error) {

	now := time.Now().UnixMicro()
	user.CreateTime = now
	user.UpdateTime = now

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.InsertUser(tx, user); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (u *UserService) ExistUsername() (bool, error) {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (u *UserService) ExistEmail() (bool, error) {
	user, err := model.GetUserByEmail(u.Email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
