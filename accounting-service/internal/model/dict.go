package model

import (
	"accounting-service/pkg/db"
	"gorm.io/gorm"
)

type Dict struct {
	Model
	LedgerId int64
	Type     int
	Name     string
	Remark   string
	HideFlag int
}

func (Dict) TableName() string {
	return "acc_dict"
}

func InsertDict(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_dict (ledger_id, type, name, remark, sort_number, hide_flag, create_time, update_time)
			select  ?, type, name, remark, sort_number, 1, ?, ? from tpl_dict where ledger_id = ?`
	if err := tx.Exec(sql, ledgerId, now, now, tLedgerId).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDictByLedgerId(tx *gorm.DB, ledgerId int64) error {
	if err := db.DB.Delete(Dict{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}
