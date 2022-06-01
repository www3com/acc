package service

import (
	"acc/internal/model"
	"acc/internal/pkg/db"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Ledger struct {
	LedgerId int64  `form:"ledgerId" binding:"required"`
	Name     string `form:"name"`
	OwnerId  int64  `form:"ownerId"`
}

func CreateLedger(tLedgerId int64, tLedgerName string, ownerId int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		return newLedger(tx, tLedgerId, tLedgerName, ownerId)
	})
	return err
}

func newLedger(tx *gorm.DB, tLedgerId int64, tLedgerName string, ownerId int64) error {
	now := time.Now().UnixMilli()
	ledger := model.Ledger{
		OwnerId:     ownerId,
		TplLedgerId: tLedgerId,
		Name:        tLedgerName,
		SortNumber:  1,
		CreateTime:  now,
		UpdateTime:  now,
	}

	ledgerId, err := model.InsertLedger(tx, &ledger)
	if err != nil {
		return errors.New(fmt.Sprintf("create ledger error, details: %s", err))
	}

	if err = createAccount(tx, tLedgerId, ledgerId, now); err != nil {
		return errors.New(fmt.Sprintf("create account error, details: %s", err))
	}

	if err = model.InsertProject(tx, ledgerId, tLedgerId, now); err != nil {
		return errors.New(fmt.Sprintf("create project error, details: %s", err))
	}

	if err = model.InsertMember(tx, ledgerId, tLedgerId, now); err != nil {
		return errors.New(fmt.Sprintf("create member error, details: %s", err))
	}

	if err = model.InsertSupplier(tx, ledgerId, tLedgerId, now); err != nil {
		return errors.New(fmt.Sprintf("create supplier error, details: %s", err))
	}
	return nil
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
			return errors.New(fmt.Sprintf("Delete ledger error, details:  %s", err))
		}

		if err := model.DeleteAccount(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete account error, details:  %s", err))
		}

		if err := model.DeleteMember(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete member error, details:  %s", err))
		}

		if err := model.DeleteSupplier(tx, ledgerId); err != nil {
			return errors.New(fmt.Sprintf("Delete supplier error, details:  %s", err))
		}

		return nil
	})
	return err
}

func ListLedger(ownerId int64) (*[]model.Ledger, error) {
	ledgers, err := model.ListLedgers(ownerId)
	if err != nil {
		return nil, err
	}

	return ledgers, nil
}

func createAccount(tx *gorm.DB, tLedgerId, ledgerId int64, now int64) error {
	tplAccounts, err := model.ListTplAccounts(tLedgerId)
	if err != nil {
		return err
	}

	for _, tplAccount := range tplAccounts {
		if tplAccount.ParentId != 0 {
			continue
		}

		parent := toAccount(ledgerId, now, 0, tplAccount)
		if err := model.InsertAccount(tx, parent); err != nil {
			return err
		}

		var children []model.Account
		for _, tplItem := range tplAccounts {
			if tplItem.ParentId == tplAccount.ID {
				account := toAccount(ledgerId, now, parent.ID, tplItem)
				children = append(children, *account)
			}
		}

		if err := model.InsertAccounts(tx, &children); err != nil {
			return err
		}
	}
	return nil
}

func toAccount(ledgerId int64, now int64, parentId int64, tplAccount *model.TplAccount) *model.Account {
	return &model.Account{
		CreateTime: now,
		UpdateTime: now,
		LedgerId:   ledgerId,
		Type:       tplAccount.Type,
		Name:       tplAccount.Name,
		Debit:      decimal.Decimal{},
		Credit:     decimal.Decimal{},
		Balance:    decimal.Decimal{},
		ParentId:   parentId,
		Icon:       tplAccount.Icon,
		CurrencyId: tplAccount.CurrencyId,
		InAsset:    tplAccount.InAsset,
		Remark:     tplAccount.Remark,
		SortNumber: tplAccount.SortNumber,
	}
}
