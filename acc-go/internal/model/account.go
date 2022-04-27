package model

import (
	"accounting-service/pkg/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Type       int
	Name       string
	Debit      decimal.Decimal
	Credit     decimal.Decimal
	Balance    decimal.Decimal
	ParentId   int64
	Icon       string
	CurrencyId int
	InAsset    int
	sortNumber int
	Remark     string
}

func (Account) TableName() string {
	return "acc_account"
}

func InsertAccount(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_account (ledger_id, type, name, debit, credit, balance, parent_id, icon, currency_id, in_asset,
                         sort_number, remark, create_time, update_time)
			select ?, type, name, 0, 0, 0, parent_id, icon, currency_id, in_asset, sort_number, remark, ?, ?
			from tpl_account where ledger_id = ?`
	if err := tx.Exec(sql, ledgerId, now, now, tLedgerId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAccount(tx *gorm.DB, accountId int64, name string) error {
	sql := "update acc_account set name = ? where id = ?"
	if err := tx.Exec(sql, name, accountId).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAccountByLedgerId(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Account{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListAccount(ledgerId int64) ([]*Account, error) {
	var accounts []*Account
	err := db.DB.Model(Account{}).Where("ledger_id = ?", ledgerId).Find(accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return accounts, nil
}
