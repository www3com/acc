package service

import (
	"acc/internal/dao"
)

type MemberService struct{}

func (s *MemberService) List(ledgerId int64) ([]*dao.Member, error) {
	return memberDao.List(ledgerId)
}
