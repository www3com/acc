package model

import (
	"acc/internal/dao"
	"github.com/shopspring/decimal"
)

type AccountOverview struct {
	Total     decimal.Decimal `json:"total"`
	Debt      decimal.Decimal `json:"debt"`
	NetAmount decimal.Decimal `json:"netAmount"`
	Details   []*dao.Account  `json:"details"`
}

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

type UpdateAccountBalanceBO struct {
	ledgerId            int64
	IncreaseAccountType int
	IncreaseAccountId   int64
	DecreaseAccountType int
	DecreaseAccountId   int64
	amount              decimal.Decimal
}
