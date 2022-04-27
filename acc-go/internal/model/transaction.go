package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              int64 `gorm:"primary_key"`
	CreateTime      int64 `json:"-"`
	UpdateTime      int64 `json:"-"`
	LedgerId        int64
	Type            int
	DebitAccountId  int64
	CreditAccountId int64
	Amount          decimal.Decimal
	Remark          string
	ProjectId       int64
	MemberId        int64
	SupplierId      int64
	Recorder        int64
}

func (Transaction) TableName() string {
	return "acc_transaction"
}

func InsertTransaction(tx *gorm.DB, transaction *Transaction) error {
	if err := tx.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
