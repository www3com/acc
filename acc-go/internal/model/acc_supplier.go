package model

import (
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

func (d *SupplierDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) (err error) {
	sql := `insert into acc_supplier (ledger_id, name, remark,  hide_flag, create_time, update_time)
			select  ?, name, remark, 0, ?, ? from tpl_supplier where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}
