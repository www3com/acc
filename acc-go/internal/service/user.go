package service

import (
	"acc/internal/consts"
	"acc/internal/model"
	"acc/internal/pkg/db"
	"acc/internal/pkg/e"
	"acc/internal/pkg/jwt"
	"acc/internal/pkg/md5"
	"acc/internal/pkg/r"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Username string
	Password string
	IP       string
}

func (u *UserService) SignIn() *r.Msg {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		logrus.Errorf("Get user by username, detail: %v", err)
		return &r.Msg{Code: e.ERROR}
	}

	if user.Username == "" {
		return &r.Msg{Code: e.USER_AUTH_ERROR}
	}

	if consts.IsUserFreeze(user.State) {
		return &r.Msg{Code: e.USER_FREEZE}
	}

	if consts.IsUserClose(user.State) {
		return &r.Msg{Code: e.USER_AUTH_ERROR}
	}

	password := md5.Encrypt(u.Password)
	if user.Password != password {
		return &r.Msg{Code: e.USER_AUTH_ERROR}
	}

	duration := 2 * time.Hour
	token, err := jwt.GenerateToken(user.ID, u.IP, duration)
	if err != nil {
		return &r.Msg{Code: e.USER_AUTH_ERROR}
	}

	return &r.Msg{
		Code: e.OK,
		Data: map[string]string{"token": token},
	}
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
