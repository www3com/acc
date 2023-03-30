package dao

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Project struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	CreateTime int64  `json:"-"`
	UpdateTime int64  `json:"-"`
	LedgerId   int64  `json:"-"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
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

func (d *ProjectDao) ListAll(ledgerId int64) ([]*Project, error) {
	var projects []*Project
	err := db.DB.Where("ledger_id = ?", ledgerId).Find(&projects).Error
	return projects, err
}

func (d *ProjectDao) List(ledgerId int64) ([]*Project, error) {
	var projects []*Project
	err := db.DB.Where("ledger_id = ? and is_show = 1", ledgerId).Order("id asc").Find(&projects).Error
	return projects, err
}
