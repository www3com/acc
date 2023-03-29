package model

import (
	"github.com/shopspring/decimal"
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          int64 `gorm:"primary_key"`
	CreateTime  int64 `json:"-"`
	TradingTime int64
	LedgerId    int64
	Type        int
	AccountId   int64
	CpAccountId int64
	Amount      decimal.Decimal
	Remark      string
	ProjectId   int64
	MemberId    int64
	SupplierId  int64
	UserId      int64
}

func (Transaction) TableName() string {
	return "acc_transaction"
}

type TransactionDao struct{}

func (d *TransactionDao) Insert(tx *gorm.DB, transaction *Transaction) error {
	return tx.Create(transaction).Error
}

func (d *TransactionDao) List(ledgerId int64,
	accounts []int64,
	cpAccounts []int64,
	projectId int64,
	memberId int64,
	supplierId int64,
	startTime int64,
	endTime int64) ([]*Transaction, error) {

	where := "where ledger_id = @ledgerId"
	if len(accounts) > 0 {
		where += "and account_id in @accounts"
	}
	if len(cpAccounts) > 0 {
		where += "and (account_id in @cpAccounts or cp_account_id in @cpAccounts)"
	}
	if projectId != 0 {
		where += "and project_id = @projectId"
	}
	if memberId != 0 {
		where += "and member_id = @memberId"
	}
	if supplierId != 0 {
		where += "and supplier_id = @supplierId"
	}
	if startTime != 0 {
		where += "and trading_time >= @startTime"
	}
	if endTime != 0 {
		where += "and trading_time <= @endTime"
	}

	if len(where) > 0 {
		where = "where " + where[4:]
	}

	sql := `select * from acc_transaction ` + where + ` order by trading_time desc `

	p := map[string]interface{}{
		"ledgerId":   ledgerId,
		"projectId":  projectId,
		"memberId":   memberId,
		"supplierId": supplierId,
		"startTime":  startTime,
		"endTime":    endTime,
		"accounts":   accounts,
		"cpAccounts": cpAccounts,
	}
	var transactions []*Transaction
	err := db.DB.Where(sql, p).Find(&transactions).Error
	return transactions, err
}
