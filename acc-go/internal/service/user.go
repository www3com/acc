package service

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/db"
	"acc/internal/pkg/e"
	"acc/internal/pkg/jwt"
	"acc/internal/pkg/md5"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Username string
	Password string
	IP       string
}

func (u *UserService) SignIn() (string, error) {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return "", e.WithStack(e.ERROR, err)
	}

	if user == nil {
		return "", e.New(e.USER_AUTH_ERROR)
	}

	if consts.IsUserFreeze(user.State) {
		return "", e.New(e.USER_FREEZE)
	}

	if consts.IsUserClose(user.State) {
		return "", e.New(e.USER_AUTH_ERROR)
	}

	password := md5.Encrypt(u.Password)
	if user.Password != password {
		return "", e.New(e.USER_AUTH_ERROR)
	}

	duration := 2 * time.Hour
	token, err := jwt.GenerateToken(user.ID, u.IP, duration)
	if err != nil {
		return "", e.New(e.USER_AUTH_ERROR)
	}

	return token, nil
}

func (u *UserService) SignUp(userId int64, user *model.User) (int64, error) {

	now := time.Now().UnixMicro()
	user.CreateTime = now
	user.UpdateTime = now
	user.Password = md5.Encrypt(user.Password)

	if err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.InsertUser(tx, user); err != nil {
			return err
		}
		return newLedger(tx, 1, "标准账本", userId)
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
