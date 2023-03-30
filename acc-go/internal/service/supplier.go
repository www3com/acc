package service

import (
	"acc/internal/dao"
)

type SupplierService struct{}

func (s *SupplierService) List(ledgerId int64) ([]*dao.Supplier, error) {
	return supplierDao.List(ledgerId)
}
