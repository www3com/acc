package service

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/db"
	"acc/internal/pkg/e"
	"acc/internal/pkg/jwt"
	"acc/internal/pkg/md5"
	"github.com/sirupsen/logrus"
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
		logrus.Errorf("get user by username failed: %v", err)
		return "", e.Error
	}

	if user.Username == "" {
		return "", e.UserAuthFailedError
	}

	if consts.IsUserFreeze(user.State) {
		return "", e.UserFreezeError
	}

	if consts.IsUserClose(user.State) {
		return "", e.UserAuthFailedError
	}

	password := md5.Encrypt(u.Password)
	if user.Password != password {
		return "", e.UserAuthFailedError
	}

	duration := 2 * time.Hour
	token, err := jwt.GenerateToken(user.ID, u.IP, duration)
	if err != nil {
		return "", e.UserAuthFailedError
	}

	return token, nil
}

func (u *UserService) SignUp(userId int64, user *model.User) (int64, error) {

	now := time.Now().UnixMicro()
	user.CreateTime = now
	user.UpdateTime = now
	user.Password = md5.Encrypt(user.Password)

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.InsertUser(tx, user); err != nil {
			return err
		}
		return newLedger(tx, 1, "标准账本", userId)
	})

	if err != nil {
		logrus.Errorf("insert user failed: %s", err)
		return 0, e.Error
	}

	return user.ID, nil
}

func (u *UserService) ExistUsername() error {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		logrus.Errorf("exist username failed: %s", err)
		return e.Error
	}

	if user.Username == "" {
		return e.UserNoUsernameError
	}

	return nil
}
