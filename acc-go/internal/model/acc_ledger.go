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
}

func (Ledger) TableName() string {
	return "acc_ledger"
}

type LedgerDao struct{}

func (d *LedgerDao) Insert(tx *gorm.DB, ledger *Ledger) (int64, error) {
	if err := tx.Create(ledger).Error; err != nil {
		return 0, err
	}
	return ledger.ID, nil
}

func (d *LedgerDao) List(ownerId int64) (ledgers *[]Ledger, err error) {
	err = db.DB.Where("owner_id = ?", ownerId).Find(&ledgers).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ledgers, nil
}
