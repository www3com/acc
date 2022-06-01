package model

import (
	"acc/internal/pkg/db"
	"gorm.io/gorm"
)

type Project struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
	LedgerId   int64
	Type       int
	Name       string
	Remark     string
	HideFlag   int
}

func (Project) TableName() string {
	return "acc_project"
}

func InsertProject(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) error {
	sql := `insert into acc_project (ledger_id,  name, remark, sort_number, hide_flag, create_time, update_time)
			select  ?, name, remark, sort_number, 1, ?, ? from tpl_project where ledger_id = ?`
	if err := tx.Exec(sql, ledgerId, now, now, tLedgerId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject(tx *gorm.DB, project *Project) error {
	sql := "update acc_project set name = ? where id = ?"
	if err := tx.Exec(sql, project.Name, project.ID).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProject(tx *gorm.DB, ledgerId int64) error {
	if err := tx.Delete(Project{}, "ledger_id = ?", ledgerId).Error; err != nil {
		return err
	}
	return nil
}

func ListProjects(ledgerId int64) ([]*Project, error) {
	var projects []*Project
	err := db.DB.Model(Project{}).Where("ledger_id = ?", ledgerId).Find(projects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return projects, nil
}
