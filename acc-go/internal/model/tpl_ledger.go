package model

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type TplLedger struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Remark string `json:"remark"`
}

func (TplLedger) TableName() string {
	return "tpl_ledger"
}

type TplLedgerDao struct{}

func (d *TplLedgerDao) Get(id int64) (*TplLedger, error) {
	var ledger TplLedger
	err := db.DB.Where("id = ?", id).First(&ledger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &ledger, nil
}
