package service

import (
	"acc/internal/model"
	"acc/internal/pkg/db"
	"acc/internal/pkg/e"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type Ledger struct {
	LedgerId int64  `form:"ledgerId" binding:"required"`
	Name     string `form:"name"`
	OwnerId  int64  `form:"ownerId"`
}

func CreateLedger(tLedgerId int64, name string, icon string, ownerId int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		return newLedger(tx, tLedgerId, name, icon, ownerId)
	})
	return err
}

func newLedger(tx *gorm.DB, tLedgerId int64, name string, icon string, ownerId int64) error {
	now := time.Now().UnixMilli()
	ledger := model.Ledger{
		OwnerId:     ownerId,
		TplLedgerId: tLedgerId,
		Name:        name,
		Icon:        icon,
		SortNumber:  1,
		CreateTime:  now,
		UpdateTime:  now,
	}

	ledgerId, err := model.InsertLedger(tx, &ledger)
	if err != nil {
		logrus.Errorf("create ledger failed: %s", err)
		return e.Error
	}

	if err = initAccount(tx, tLedgerId, ledgerId, now); err != nil {
		logrus.Errorf("create account failed: %s", err)
		return e.Error
	}

	if err = model.InsertProject(tx, ledgerId, tLedgerId, now); err != nil {
		logrus.Errorf("create project failed: %s", err)
		return e.Error
	}

	if err = model.InsertMember(tx, ledgerId, tLedgerId, now); err != nil {
		logrus.Errorf("create member failed: %s", err)
		return e.Error
	}

	if err = model.InsertSupplier(tx, ledgerId, tLedgerId, now); err != nil {
		logrus.Errorf("create supplier failed: %s", err)
		return e.Error
	}

	return nil
}

func UpdateLedger(ledgerId int64, name string) error {
	ledger := model.Ledger{ID: ledgerId, Name: name}
	if err := model.UpdateLedger(&ledger); err != nil {
		logrus.Errorf("update ledger failed: %s", err)
		return e.Error
	}
	return nil
}

func DeleteLedger(ledgerId int64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteLedger(tx, ledgerId); err != nil {
			logrus.Errorf("delete ledger failed: %s", err)
			return e.Error
		}

		if err := model.DeleteAccount(tx, ledgerId); err != nil {
			logrus.Errorf("delete account failed: %s", err)
			return e.Error
		}

		if err := model.DeleteMember(tx, ledgerId); err != nil {
			logrus.Errorf("delete member failed: %s", err)
			return e.Error
		}

		if err := model.DeleteSupplier(tx, ledgerId); err != nil {
			logrus.Errorf("delete supplier failed: %s", err)
			return e.Error
		}

		return nil
	})
	return err
}

func ListLedgers(ownerId int64) (*[]model.Ledger, error) {
	ledgers, err := model.ListLedgers(ownerId)
	if err != nil {
		logrus.Errorf("list ledger failed: %s", err)
		return nil, e.Error
	}

	return ledgers, nil
}

func initAccount(tx *gorm.DB, tLedgerId, ledgerId int64, now int64) error {
	tplAccounts, err := model.ListTplAccounts(tLedgerId)
	if err != nil {
		logrus.Errorf("list template account failed: %s", err)
		return e.Error
	}

	for _, tplAccount := range tplAccounts {
		if tplAccount.ParentId != 0 {
			continue
		}

		parent := toAccount(ledgerId, now, 0, tplAccount)
		if err := model.InsertAccount(tx, parent); err != nil {
			logrus.Errorf("insert root account failed: %s", err)
			return e.Error
		}

		var children []model.Account
		for _, tplItem := range tplAccounts {
			if tplItem.ParentId == tplAccount.ID {
				account := toAccount(ledgerId, now, parent.ID, tplItem)
				children = append(children, *account)
			}
		}

		if err := model.InsertAccounts(tx, &children); err != nil {
			logrus.Errorf("batch insert account failed: %s", err)
			return e.Error
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
