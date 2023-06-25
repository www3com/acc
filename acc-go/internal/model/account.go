package model

import (
	"github.com/shopspring/decimal"
)

type AccountOverview struct {
	Total     decimal.Decimal `json:"total"`
	Debt      decimal.Decimal `json:"debt"`
	NetAmount decimal.Decimal `json:"netAmount"`
	Details   []*AccountVO    `json:"details"`
}

type AccountVO struct {
	ID         int64           `json:"id" gorm:"primary_key"`
	Type       int             `json:"type"`
	Name       string          `json:"name"`
	Code       string          `json:"code"`
	Level      int             `json:"level"`
	Debit      decimal.Decimal `json:"debit"`
	Credit     decimal.Decimal `json:"credit"`
	Balance    decimal.Decimal `json:"balance"`
	Icon       string          `json:"icon"`
	CurrencyId int             `json:"currencyId"`
	Remark     string          `json:"remark"`
	Children   []*AccountVO    `json:"children" gorm:"-"`
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
