package model

import (
	"acc/internal/pkg/db"
	"gorm.io/gorm"
)

type Dict struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Type       int
	Name       string
	Remark     string
	HideFlag   int
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

func UpdateDict(tx *gorm.DB, dict *Dict) error {
	sql := "update acc_dict set name = ? where id = ?"
	if err := tx.Exec(sql, dict.Name, dict.ID).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDictByLedgerId(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Dict{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListDicts(ledgerId int64, dictType int64) ([]*Dict, error) {
	var dicts []*Dict
	err := db.DB.Model(Dict{}).Where("ledger_id = ?, type = ?", ledgerId, dictType).Find(dicts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return dicts, nil
}
