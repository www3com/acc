package service

import (
	"acc/internal/model"
)

type ProjectService struct{}

func (s *ProjectService) List(ledgerId int64) ([]*model.Project, error) {
	return projectDao.List(ledgerId)
}
