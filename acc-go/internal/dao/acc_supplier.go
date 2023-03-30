package dao

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Supplier struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	CreateTime int64  `json:"-"`
	UpdateTime int64  `json:"-"`
	LedgerId   int64  `json:"-"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
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
	err := db.DB.Where("ledger_id = ?", ledgerId).Find(&suppliers).Error
	return suppliers, err
}

func (d *SupplierDao) List(ledgerId int64) ([]*Supplier, error) {
	var suppliers []*Supplier
	err := db.DB.Where("ledger_id = ? and is_show = 1", ledgerId).Order("id asc").Find(&suppliers).Error
	return suppliers, err
}
