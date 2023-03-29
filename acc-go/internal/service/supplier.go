package service

import (
	"acc/internal/model"
)

type SupplierService struct{}

func (s *SupplierService) List(ledgerId int64) ([]*model.Supplier, error) {
	return supplierDao.List(ledgerId)
}
