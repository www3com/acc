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
	projects []int64,
	members []int64,
	suppliers []int64,
	startDate int64,
	endDate int64) ([]*Transaction, error) {

	where := "ledger_id = @ledgerId"
	if len(accounts) > 0 {
		where += "and account_id in @accounts"
	}
	if len(cpAccounts) > 0 {
		where += "and (account_id in @cpAccounts or cp_account_id in @cpAccounts)"
	}
	if len(projects) > 0 {
		where += "and project_id in @projects"
	}
	if len(members) > 0 {
		where += "and member_id = @members"
	}
	if len(suppliers) > 0 {
		where += "and supplier_id = @suppliers"
	}
	if startDate != 0 {
		where += "and trading_time >= @startDate"
	}
	if endDate != 0 {
		where += "and trading_time <= @endDate"
	}

	//sql := `select * from acc_transaction ` + where + ` order by trading_time desc `

	p := map[string]interface{}{
		"ledgerId":   ledgerId,
		"projects":   projects,
		"members":    members,
		"suppliers":  suppliers,
		"startTime":  startDate,
		"endTime":    endDate,
		"accounts":   accounts,
		"cpAccounts": cpAccounts,
	}
	var transactions []*Transaction
	err := db.DB.Where(where, p).Order("trading_time desc").Find(&transactions).Error
	return transactions, err
}
