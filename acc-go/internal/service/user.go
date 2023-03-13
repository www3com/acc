package service

import "acc/internal/model"

type UserService struct{}

var UserSrv = new(UserService)

func (s *UserService) SignIn(username, password, ip string) (string, error) {
	model.TplAccountDao.List(1)
	return "", nil
}
