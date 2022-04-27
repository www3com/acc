package model

import (
	"accounting-service/pkg/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	Model
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

func DeleteAccountByLedgerId(tx *gorm.DB, ledgerId int64) error {
	if err := db.DB.Delete(Account{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}
