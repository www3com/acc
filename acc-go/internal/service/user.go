package service

import (
	"acc/internal/consts"
	"acc/internal/model"
	"github.com/upbos/go-saber/db"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/log"
	"github.com/upbos/go-saber/md5"
	"gorm.io/gorm"
	"time"
)

type UserService struct{}

func (s *UserService) SignUp(user *model.User) error {

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
		log.Errorf(err, "insert user")
		return e.Err
	}

	return nil
}
