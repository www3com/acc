package service

import (
	"acc/internal/model"
	"acc/internal/pkg/db"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	LedgerId        int64           `form:"ledgerId" binding:"required"`
	DebitAccountId  int64           `form:"debitAccountId" binding:"required"`
	CreditAccountId int64           `form:"creditAccountId" binding:"required"`
	Amount          decimal.Decimal `form:"amount" binding:"required"`
	Type            int             `form:"type" binding:"required"`
	Remark          string          `form:"remark"`
	ProjectId       int64           `form:"projectId"`
	MemberId        int64           `form:"memberId"`
	SupplierId      int64           `form:"supplierId"`
	Recorder        int64
}

func CreateTransaction(transaction *Transaction) error {
	now := time.Now().UnixMicro()
	trans := model.Transaction{
		CreateTime:      now,
		UpdateTime:      now,
		LedgerId:        transaction.LedgerId,
		Type:            transaction.Type,
		DebitAccountId:  transaction.DebitAccountId,
		CreditAccountId: transaction.CreditAccountId,
		Amount:          transaction.Amount,
		Remark:          transaction.Remark,
		ProjectId:       transaction.ProjectId,
		MemberId:        transaction.MemberId,
		SupplierId:      transaction.SupplierId,
		Recorder:        transaction.Recorder,
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.InsertTransaction(tx, &trans); err != nil {
			return err
		} else {
			return nil
		}
	})

	return err
}
