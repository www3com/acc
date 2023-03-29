package service

import (
	"acc/internal/model"
)

type MemberService struct{}

func (s *MemberService) List(ledgerId int64) ([]*model.Member, error) {
	return memberDao.List(ledgerId)
}
