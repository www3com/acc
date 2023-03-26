package model

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Supplier struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Name       string
	Remark     string
	IsShow     bool
}

func (Supplier) TableName() string {
	return "acc_supplier"
}

type SupplierDao struct{}

func (d *SupplierDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_supplier (ledger_id, name, remark,  is_show, create_time, update_time)
			select  ?, name, remark, 1, ?, ? from tpl_supplier where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}

func (d *SupplierDao) ListAll(ledgerId int64) ([]*Supplier, error) {
	var suppliers []*Supplier
	err := db.DB.Where("ledgerId = ?", ledgerId).Find(&suppliers).Error
	return suppliers, err
}

func (d *SupplierDao) List(ledgerId int64) ([]*Supplier, error) {
	var suppliers []*Supplier
	err := db.DB.Where("ledgerId = ? and is_show = 1", ledgerId).Order("id asc").Find(&suppliers).Error
	return suppliers, err
}
