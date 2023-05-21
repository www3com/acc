package dao

import (
	"acc/internal/model"
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

func (d *TransactionDao) List(tran *model.TransactionQuery) ([]*Transaction, error) {
	var transactions []*Transaction
	err := db.DB.Where(buildWhere(tran), tran).Order("trading_time desc").Find(&transactions).Error
	return transactions, err
}

func (d *TransactionDao) ListTotal(tran *model.TransactionQuery) ([]*model.TransactionTotalBO, error) {
	var transactions []*model.TransactionTotalBO
	sql := "select" + " Type, sum(amount) Amount from acc_transaction where type in(1,2) and " + buildWhere(tran) + " group by type"
	err := db.DB.Table("acc_transaction").Raw(sql, tran).Scan(&transactions).Error
	return transactions, err
}

func buildWhere(tran *model.TransactionQuery) string {
	where := "ledger_id = @LedgerId"
	if len(tran.Accounts) > 0 {
		where += " and account_id in @Accounts"
	}
	if len(tran.CpAccounts) > 0 {
		where += " and (account_id in @CpAccounts or cp_account_id in @CpAccounts)"
	}
	if len(tran.Projects) > 0 {
		where += " and project_id in @Projects"
	}
	if len(tran.Members) > 0 {
		where += " and member_id = @Members"
	}
	if len(tran.Suppliers) > 0 {
		where += " and supplier_id = @Suppliers"
	}
	if tran.StartTime != 0 {
		where += " and trading_time >= @StartTime"
	}
	if tran.EndTime != 0 {
		where += " and trading_time <= @EndTime"
	}
	return where
}
