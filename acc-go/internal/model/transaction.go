package model

import "github.com/shopspring/decimal"

type TransactionTotalBO struct {
	Type   int ``
	Amount decimal.Decimal
}

type TransactionTotalVO struct {
	Income  string `json:"income"`
	Expense string `json:"expense"`
	Balance string `json:"balance"`
}

type TransactionBO struct {
	LedgerId    int64
	UserId      int64
	TradingTime int64           `form:"tradingTime" binding:"required"`
	Type        int             `form:"type" binding:"required,min=1,max=7"`
	AccountId   int64           `form:"accountId" binding:"required"`
	CpAccountId int64           `form:"cpAccountId" binding:"required"`
	Amount      decimal.Decimal `form:"amount" binding:"required"`
	Remark      string          `form:"remark"`
	ProjectId   int64           `form:"projectId"`
	MemberId    int64           `form:"memberId"`
	SupplierId  int64           `form:"supplierId"`
}

type TransactionVO struct {
	Id          int64  `json:"id"`
	TradingTime string `json:"tradingTime"`
	Type        string `json:"type"`
	Account     string `json:"account"`
	CpAccount   string `json:"cpAccount"`
	Project     string `json:"project"`
	Member      string `json:"member"`
	Supplier    string `json:"supplier"`
	Amount      string `json:"amount"`
	Remark      string `json:"remark"`
}

type TransactionQuery struct {
	LedgerId   int64
	Accounts   []int64 `form:"accounts"`
	CpAccounts []int64 `form:"cpAccounts"`
	Projects   []int64 `form:"projects"`
	Members    []int64 `form:"members"`
	Suppliers  []int64 `form:"suppliers"`
	StartTime  int64   `form:"startDate"`
	EndTime    int64   `form:"startDate"`
}
