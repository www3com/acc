package model

type Model struct {
	ID         int64 `gorm:"primary_key"`
	CreateTime int64
	UpdateTime int64
}
