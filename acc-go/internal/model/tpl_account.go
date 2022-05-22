package model

import (
	"accounting-service/internal/pkg/db"
	"gorm.io/gorm"
)

type TplAccount struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	LedgerId   int64  `json:"-"`
	Type       int    `json:"-"`
	Name       string `json:"name"`
	ParentId   int64  `json:"-"`
	Icon       string `json:"icon"`
	CurrencyId int    `json:"currencyId"`
	InAsset    int    `json:"InAsset"`
	SortNumber int    `json:"sortNumber"`
	Remark     string `json:"remark"`
}

func (TplAccount) TableName() string {
	return "tpl_account"
}

func ListTplAccounts(ledgerId int64) ([]*TplAccount, error) {
	var accounts []*TplAccount
	err := db.DB.Model(TplAccount{}).Where("ledger_id = ?", ledgerId).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return accounts, nil
}
