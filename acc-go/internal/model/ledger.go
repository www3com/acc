package model

import (
	"accounting-service/pkg/db"
	"gorm.io/gorm"
)

type Ledger struct {
	Model
	OwnerId     int64
	TplLedgerId int64
	Name        string
	Remark      string
	SortNumber  int
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

func UpdateLedger(ledger Ledger) error {
	if err := db.DB.Model(ledger).Where("id = ?", ledger.ID).Update("name", ledger.Name).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLedger(tx *gorm.DB, id int64) error {
	if err := db.DB.Delete(Ledger{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func ListLedger(ownerId int) ([]*Ledger, error) {
	var ledgers []*Ledger
	err := db.DB.Where("owner_id = ?", ownerId).Find(ledgers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ledgers, nil
}
