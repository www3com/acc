package service

import (
	"acc/internal/dao"
)

type ProjectService struct{}

func (s *ProjectService) List(ledgerId int64) ([]*dao.Project, error) {
	return projectDao.List(ledgerId)
}
