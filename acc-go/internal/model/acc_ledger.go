package model

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Ledger struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	CreateTime int64  `json:"-"`
	UpdateTime int64  `json:"-"`
	OwnerId    int64  `json:"-"`
	TLedgerId  int64  `json:"-"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Remark     string `json:"remark"`
	Selected   int    `json:"selected"`
}

func (Ledger) TableName() string {
	return "acc_ledger"
}

type LedgerDao struct{}

func (d *LedgerDao) SetDefault(ledgerId int64) error {
	return db.DB.Model(&Ledger{}).Where("ledger_id = ?", ledgerId).Update("selected", 1).Error
}

func (d *LedgerDao) CancelDefault(ownerId int64) error {
	return db.DB.Model(&Ledger{}).Where("owner_id = ?", ownerId).Update("selected", 0).Error
}

func (d *LedgerDao) Default(ownerId int64) (*Ledger, error) {
	var ledger Ledger
	err := db.DB.Where("owner_id = ? and selected = 1", ownerId).First(&ledger).Error
	return &ledger, err
}

func (d *LedgerDao) List(ownerId int64) (ledgers *[]Ledger, err error) {
	err = db.DB.Where("owner_id = ?", ownerId).Find(&ledgers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ledgers, nil
}

func (d *LedgerDao) Insert(tx *gorm.DB, ledger *Ledger) (int64, error) {
	if err := tx.Create(ledger).Error; err != nil {
		return 0, err
	}
	return ledger.ID, nil
}
