package model

import (
	"acc/internal/pkg/db"
	"gorm.io/gorm"
)

type Supplier struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Type       int
	Name       string
	Remark     string
	HideFlag   int
}

func (Supplier) TableName() string {
	return "acc_supplier"
}

func InsertSupplier(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_supplier (ledger_id, name, remark, sort_number, hide_flag, create_time, update_time)
			select  ?, name, remark, sort_number, 1, ?, ? from tpl_supplier where ledger_id = ?`
	if err := tx.Exec(sql, ledgerId, now, now, tLedgerId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSupplier(tx *gorm.DB, supplier *Supplier) error {
	sql := "update acc_supplier set name = ? where id = ?"
	if err := tx.Exec(sql, supplier.Name, supplier.ID).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSupplier(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Supplier{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListSuppliers(ledgerId int64) ([]*Supplier, error) {
	var suppliers []*Supplier
	err := db.DB.Model(Member{}).Where("ledger_id = ?", ledgerId).Find(suppliers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return suppliers, nil
}
