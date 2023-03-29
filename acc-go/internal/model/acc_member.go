package model

import (
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Member struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	CreateTime int64  `json:"-"`
	UpdateTime int64  `json:"-"`
	LedgerId   int64  `json:"-"`
	Name       string `json:"name"`
	Remark     string `json:"remark"`
}

func (Member) TableName() string {
	return "acc_member"
}

type MemberDao struct{}

func (d *MemberDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_member (ledger_id, name, remark, is_show, create_time, update_time)
			select  ?, name, remark, 1, ?, ? from tpl_member where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}

func (d *MemberDao) ListAll(ledgerId int64) ([]*Member, error) {
	var members []*Member
	err := db.DB.Where("ledger_id = ?", ledgerId).Find(&members).Error
	return members, err
}

func (d *MemberDao) List(ledgerId int64) ([]*Member, error) {
	var members []*Member
	err := db.DB.Where("ledger_id = ? and is_show = 1", ledgerId).Order("id asc").Find(&members).Error
	return members, err
}
