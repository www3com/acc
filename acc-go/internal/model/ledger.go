package model

import (
	"accounting-service/pkg/db"
	"gorm.io/gorm"
)

type Ledger struct {
	ID          int64 `gorm:"primary_key"`
	CreateTime  int64 `json:"-"`
	UpdateTime  int64 `json:"-"`
	OwnerId     int64 `json:"-"`
	TplLedgerId int64 `json:"-"`
	Name        string
	Remark      string
	SortNumber  int `json:"-"`
}

func (Ledger) TableName() string {
	return "acc_ledger"
}

func InsertLedger(tx *gorm.DB, ledger *Ledger) (int64, error) {
	if err := tx.Create(ledger).Error; err != nil {
		return 0, err
	}
	return ledger.ID, nil
}

func UpdateLedger(ledger *Ledger) error {
	if err := db.DB.Model(ledger).Where("id = ?", ledger.ID).Update("name", ledger.Name).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLedger(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Ledger{}, "id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListLedger(ownerId int64) (*[]Ledger, error) {
	var ledgers *[]Ledger
	err := db.DB.Where("owner_id = ?", ownerId).Find(&ledgers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ledgers, nil
}
