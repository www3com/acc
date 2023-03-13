package model

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type TplAccount struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	LedgerId   int64  `json:"-"`
	Type       int    `json:"-"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Icon       string `json:"icon"`
	CurrencyId int    `json:"currencyId"`
	Leaf       int    `json:"leaf"`
	Remark     string `json:"remark"`
}

func (TplAccount) TableName() string {
	return "tpl_account"
}

type TplAccountModel struct {
	db *gorm.DB
}

var TplAccountDao = TplAccountModel{db: db.DB}

func (d *TplAccountModel) List(ledgeId int64) (accounts []*TplAccount, err error) {
	err = d.db.Model(TplAccount{}).Where("ledger_id = ?", ledgeId).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return accounts, nil
}
