package service

import (
	"accounting-service/internal/model"
	"accounting-service/pkg/db"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Ledger struct {
	LedgerId int64  `form:"ledgerId" binding:"required"`
	Name     string `form:"name"`
	OwnerId  int64  `form:"ownerId"`
}

func CreateLedger(tLedgerId int64, ownerId int64) error {
	now := time.Now().UnixMilli()
	ledger := model.Ledger{
		OwnerId:     ownerId,
		TplLedgerId: tLedgerId,
		Name:        "标准账本",
		SortNumber:  1,
		CreateTime:  now,
		UpdateTime:  now,
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		ledgerId, err := model.InsertLedger(tx, &ledger)
		if err != nil {
			return errors.New(fmt.Sprintf("create ledger error, template ledger id: %d, details: %s", tLedgerId, err))
		}

		if err = model.InsertAccount(tx, ledgerId, tLedgerId, now); err != nil {
			return errors.New(fmt.Sprintf("create account error, template ledger id: %d, details: %s", tLedgerId, err))
		}

		if err = model.InsertDict(tx, ledgerId, tLedgerId, now); err != nil {
			return errors.New(fmt.Sprintf("create dict error, template ledger id: %d, details: %s", tLedgerId, err))
		}
		return nil
	})
	return err
}

func UpdateLedger(ledgerId int64, name string) error {
	ledger := model.Ledger{ID: ledgerId, Name: name}
	if err := model.UpdateLedger(&ledger); err != nil {
		return err
	}
	return nil
}

func DeleteLedger(ledgerId int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteLedger(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete ledger error, ledger id: %d, details:  %s", ledgerId, err))
		}

		if err := model.DeleteAccountByLedgerId(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete account error, ledger id: %d, details:  %s", ledgerId, err))
		}

		if err := model.DeleteDictByLedgerId(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete dict error, ledger id: %d, details:  %s", ledgerId, err))
		}

		return nil
	})
	return err
}

func ListLedger(ownerId int64) (*[]model.Ledger, error) {
	ledgers, err := model.ListLedger(ownerId)
	if err != nil {
		return nil, err
	}
	return ledgers, nil
}
