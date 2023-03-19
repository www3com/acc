package model

import (
	"github.com/shopspring/decimal"
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Account struct {
	ID         int64           `json:"id" gorm:"primary_key"`
	CreateTime int64           `json:"-"`
	UpdateTime int64           `json:"-"`
	LedgerId   int64           `json:"-"`
	Type       int             `json:"-"`
	Name       string          `json:"name"`
	Code       string          `json:"code"`
	Debit      decimal.Decimal `json:"debit"`
	Credit     decimal.Decimal `json:"credit"`
	Balance    decimal.Decimal `json:"balance"`
	Icon       string          `json:"icon"`
	CurrencyId int             `json:"currencyId"`
	Remark     string          `json:"remark"`
	Children   []*Account      `json:"children" gorm:"-"`
}

func (Account) TableName() string {
	return "acc_account"
}

type AccountDao struct{}

func (d *AccountDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) (err error) {
	sql := `insert into acc_account (ledger_id, type, name, code, debit, credit, balance, icon, currency_id, remark, create_time, update_time)
			select  ?, type, name, code, 0, 0, 0, icon, currency_id, remark, ?, ? from tpl_account where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}

func (d *AccountDao) List(ledgerId int64) (accounts []*Account, err error) {
	err = db.DB.Model(Account{}).Where("ledger_id = ? and type in (1,2)", ledgerId).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return accounts, nil
}
