package service

import (
	"accounting-service/internal/misc/consts"
	"accounting-service/internal/model"
	"accounting-service/internal/pkg/db"
	"accounting-service/internal/pkg/jwt"
	"accounting-service/internal/pkg/md5"
	"accounting-service/internal/pkg/r"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Username string
	Password string
	IP       string
}

func (u *UserService) SignIn() (*r.R[string], error) {
	user, err := model.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	if user.Username == "" {
		return r.Fail[string](consts.USER_SIGN_IN_ERROR), nil
	}

	if consts.IsUserFreeze(user.State) {
		return r.Fail[string](consts.USER_SIGN_IN_FREEZE), nil
	}

	if consts.IsUserClose(user.State) {
		return r.Fail[string](consts.USER_SIGN_IN_ERROR), nil
	}

	password := md5.Encrypt(u.Password)
	if user.Password != password {
		return r.Fail[string](consts.USER_SIGN_IN_ERROR), nil
	}

	duration := 2 * time.Hour
	token, err := jwt.GenerateToken(user.ID, u.IP, duration)
	if err != nil {
		return r.Fail[string](consts.USER_SIGN_IN_ERROR), nil
	}

	return r.Ok[string](token), nil

}

func (u *UserService) SignUp(user *model.User) (int64, error) {

	now := time.Now().UnixMicro()
	user.CreateTime = now
	user.UpdateTime = now
	user.Password = md5.Encrypt(user.Password)

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
