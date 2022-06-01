package model

import (
	"acc/internal/pkg/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	ID         int64           `json:"id" gorm:"primary_key"`
	CreateTime int64           `json:"-"`
	UpdateTime int64           `json:"-"`
	LedgerId   int64           `json:"-"`
	Type       int             `json:"-"`
	Name       string          `json:"name"`
	Debit      decimal.Decimal `json:"debit"`
	Credit     decimal.Decimal `json:"credit"`
	Balance    decimal.Decimal `json:"balance"`
	ParentId   int64           `json:"-"`
	Icon       string          `json:"icon"`
	CurrencyId int             `json:"currencyId"`
	InAsset    int             `json:"InAsset"`
	SortNumber int             `json:"sortNumber"`
	Remark     string          `json:"remark"`
	Children   []*Account      `json:"children" gorm:"-"`
}

func (Account) TableName() string {
	return "acc_account"
}

func InsertAccount(tx *gorm.DB, account *Account) error {
	if err := tx.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func InsertAccounts(tx *gorm.DB, accounts *[]Account) error {
	if err := tx.CreateInBatches(accounts, len(*accounts)).Error; err != nil {
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

func DeleteAccount(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Account{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListAccounts(ledgerId int64) ([]*Account, error) {
	var accounts []*Account
	err := db.DB.Model(Account{}).Where("ledger_id = ?", ledgerId).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return accounts, nil
}
