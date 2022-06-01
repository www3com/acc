package model

import (
	"acc/internal/pkg/db"
	"gorm.io/gorm"
)

type Member struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Type       int
	Name       string
	Remark     string
	HideFlag   int
}

func (Member) TableName() string {
	return "acc_member"
}

func InsertMember(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_member (ledger_id, name, remark, sort_number, hide_flag, create_time, update_time)
			select  ?, name, remark, sort_number, 1, ?, ? from tpl_member where ledger_id = ?`
	if err := tx.Exec(sql, ledgerId, now, now, tLedgerId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMember(tx *gorm.DB, member *Member) error {
	sql := "update acc_member set name = ? where id = ?"
	if err := tx.Exec(sql, member.Name, member.ID).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMember(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Member{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListMembers(ledgerId int64) ([]*Member, error) {
	var members []*Member
	err := db.DB.Model(Member{}).Where("ledger_id = ?", ledgerId).Find(members).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return members, nil
}
