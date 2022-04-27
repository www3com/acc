package service

import (
	"accounting-service/internal/model"
	"accounting-service/pkg/db"
	"accounting-service/pkg/logger"
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
			logger.Error("create ledger error, template ledger id: {}, details: ", tLedgerId, err)
			return err
		}

		if err = model.InsertAccount(tx, ledgerId, tLedgerId, now); err != nil {
			logger.Error("create account error, template ledger id: {}, details: ", ledgerId, err)
			return err
		}

		if err = model.InsertDict(tx, ledgerId, tLedgerId, now); err != nil {
			logger.Error("create dict error, template ledger id: {}, details: ", ledgerId, err)
			return err
		}

		return nil
	})
	return err
}

func UpdateLedger(ledgerId int64, name string) error {
	ledger := model.Ledger{ID: ledgerId, Name: name}
	if err := model.UpdateLedger(&ledger); err != nil {
		logger.Error("Update Ledger name error, ledger id: {}, details: ", ledgerId, err)
		return err
	}
	return nil
}

func DeleteLedger(ledgerId int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteLedger(tx, ledgerId); err != nil {
			logger.Error("Delete ledger error, ledger id: {}, details: ", ledgerId, err)
			return err
		}

		if err := model.DeleteAccountByLedgerId(tx, ledgerId); err != nil {
			logger.Error("Delete account error, ledger id: {}, details: ", ledgerId, err)
			return err
		}

		if err := model.DeleteDictByLedgerId(tx, ledgerId); err != nil {
			logger.Error("Delete dict error, ledger id: {}, details: ", ledgerId, err)
			return err
		}

		return nil
	})
	return err
}

func ListLedger(ownerId int64) (*[]model.Ledger, error) {
	ledgers, err := model.ListLedger(ownerId)
	if err != nil {
		logger.Error("Query user account error, user id: {}, details: ", ownerId, err)
		return nil, err
	}
	return ledgers, nil
}
