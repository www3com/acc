package model

import (
	"gorm.io/gorm"
)

type Project struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Name       string
	Remark     string
	IsShow     bool
}

func (Project) TableName() string {
	return "acc_project"
}

type ProjectDao struct{}

func (d *ProjectDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_project (ledger_id,  name, remark, is_show, create_time, update_time)
			select  ?, name, remark, 1, ?, ? from tpl_project where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}
