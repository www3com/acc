package model

import (
	"acc/internal/service"
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

func (d *TransactionDao) List(ledgerId int64, tran *service.TransactionBO) ([]*Transaction, error) {

	where := "where ledger_id = @ledgerId"
	if tran.ProjectId != 0 {
		where += "and project_id = @projectId"
	}
	if tran.MemberId != 0 {
		where += "and member_id = @memberId"
	}
	if tran.SupplierId != 0 {
		where += "and supplier_id = @supplierId"
	}
	if tran.StartTime != 0 {
		where += "and trading_time >= @startTime"
	}
	if tran.EndTime != 0 {
		where += "and trading_time <= @endTime"
	}

	if len(tran.Accounts) > 0 {
		where += "and account_id in @accounts"
	}
	if len(tran.CpAccounts) > 0 {
		where += "and (account_id in @cpAccounts or cp_account_id in @cpAccounts)"
	}

	if len(where) > 0 {
		where = "where " + where[4:]
	}

	sql := `select * from acc_transaction ` + where + ` order by trading_time desc `

	p := map[string]interface{}{
		"ledgerId":   ledgerId,
		"projectId":  tran.ProjectId,
		"memberId":   tran.MemberId,
		"supplierId": tran.SupplierId,
		"startTime":  tran.StartTime,
		"endTime":    tran.EndTime,
		"accounts":   tran.Accounts,
		"cpAccounts": tran.CpAccounts,
	}
	var transactions []*Transaction
	err := db.DB.Where(sql, p).Find(&transactions).Error
	return transactions, err
}
