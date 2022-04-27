package service

import (
	"accounting-service/internal/model"
	"accounting-service/pkg/db"
	"accounting-service/pkg/logger"
	"gorm.io/gorm"
	"time"
)

func CreateLedger(tLedgerId int64, ownerId int64) error {
	now := time.Now().UnixMilli()
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		ledger := model.Ledger{
			Model:       model.Model{CreateTime: now, UpdateTime: now},
			OwnerId:     ownerId,
			TplLedgerId: tLedgerId,
			Name:        "标准账本",
			SortNumber:  1,
		}
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

func DeleteLedger(id int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteLedger(tx, id); err != nil {
			logger.Error("Delete ledger error, ledger id: {}, details: ", id, err)
			return err
		}

		if err := model.DeleteAccountByLedgerId(tx, id); err != nil {
			logger.Error("Delete account error, ledger id: {}, details: ", id, err)
			return err
		}

		if err := model.DeleteDictByLedgerId(tx, id); err != nil {
			logger.Error("Delete dict error, ledger id: {}, details: ", id, err)
			return err
		}

		return nil
	})
	return err
}
