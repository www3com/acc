package model

import (
	"github.com/shopspring/decimal"
)

type AccountBO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Remark   string `json:"remark"`
	ParentId int64  `json:"parentId"`
}

type UpdateAccountBO struct {
	Type   string          `json:"type"`
	Id     int64           `form:"id" binding:"required"`
	Name   string          `form:"name"`
	Remark string          `form:"remark"`
	Amount decimal.Decimal `form:"amount"`
}

type DeleteAccountBO struct {
	LedgerId int64
	Code     string `form:"code"`
}
