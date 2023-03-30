package service

import (
	"acc/internal/consts"
	"acc/internal/consts/userstate"
	"acc/internal/dao"
	"errors"
	"github.com/upbos/go-saber/db"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/jwt"
	"github.com/upbos/go-saber/log"
	"github.com/upbos/go-saber/md5"
	"github.com/upbos/go-saber/uid"
	"gorm.io/gorm"
	"time"
)

type UserService struct{}

// SignIn 用户登录
func (s *UserService) SignIn(username, password, ip string) (string, error) {
	user, err := userDao.Get(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return consts.Empty, consts.ErrUserAuthFailed
		}
		log.Errorf(err, "get userstate by username failed")
		return consts.Empty, e.Err
	}

	if userstate.IsFreeze(user.State) {
		return consts.Empty, consts.ErrUserFreeze
	}

	if userstate.IsClose(user.State) {
		return consts.Empty, consts.ErrUserAuthFailed
	}

	pwd := md5.Encrypt(password)
	if user.Password != pwd {
		return consts.Empty, consts.ErrUserAuthFailed
	}

	duration := 2 * time.Hour

	token, err := jwt.GenerateToken(uid.Uid(user.ID), ip, duration)
	if err != nil {
		return consts.Empty, consts.ErrUserAuthFailed
	}

	return token, nil
}

// SignUp 用户注册
func (s *UserService) SignUp(user *dao.User) error {

	user.CreateTime = time.Now().UnixMicro()
	user.UpdateTime = user.CreateTime
	user.Password = md5.Encrypt(user.Password)
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := userDao.Insert(tx, user); err != nil {
			return err
		}
		return ledgerService.New(tx, consts.RootLedgerId, user.ID)
	})

	if err != nil {
		log.Error(err, "insert userstate")
		return e.Err
	}

	return nil
}

func (s *UserService) Exist(username string) (bool, error) {
	_, err := userDao.Get(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
