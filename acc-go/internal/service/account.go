package service

import (
	acc "acc/internal/consts/account"
	"acc/internal/model"
	"acc/internal/pkg/code"
	"github.com/shopspring/decimal"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/log"
	"time"
)

type Account struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Remark   string `json:"remark"`
	ParentId int64  `json:"parentId"`
}

type AccountDetails struct {
	Total     decimal.Decimal  `json:"total"`
	Debt      decimal.Decimal  `json:"debt"`
	NetAmount decimal.Decimal  `json:"netAmount"`
	Details   []*model.Account `json:"details"`
}

type AccountService struct{}

func (s *AccountService) Create(ledgerId int64, account *Account) error {
	parent, err := accountDao.Get(ledgerId, account.ParentId)
	if err != nil {
		log.Error(err, "get parent account")
		return err
	}

	maxCode, err := accountDao.MaxCode(ledgerId, parent.Code, code.Level(parent.Code)+1)
	if err != nil {
		log.Error(err, "get account child max code")
		return err
	}

	now := time.Now().UnixMicro()
	nextCode, err := code.Next(maxCode)
	if nextCode == "" {
		nextCode = code.FirstChild(parent.Code)
	}
	if err != nil {
		return err
	}
	dbAccount := model.Account{
		CreateTime: now,
		UpdateTime: now,
		LedgerId:   ledgerId,
		Type:       parent.Type,
		Name:       account.Name,
		Code:       nextCode,
		Level:      code.Level(nextCode),
		Icon:       parent.Icon,
		CurrencyId: parent.CurrencyId,
		Remark:     account.Remark,
	}

	return accountDao.Insert(&dbAccount)
}

func (s *AccountService) Update(ledgerId int64, account *Account) error {
	now := time.Now().UnixMicro()
	dbAccount := model.Account{
		UpdateTime: now,
		LedgerId:   ledgerId,
		ID:         account.Id,
		Name:       account.Name,
		Remark:     account.Remark,
	}

	return accountDao.Update(&dbAccount)
}

func (s *AccountService) List(ledgerId int64) (*AccountDetails, error) {

	accounts, err := accountDao.List(ledgerId)
	if err != nil {
		return nil, e.Wrap("Query account error", err)
	}

	var accountMap = make(map[string]*model.Account, len(accounts))
	for _, account := range accounts {
		accountMap[account.Code] = account
	}

	// 构造树
	var calcAccounts []*model.Account
	for _, account := range accounts {
		if code.Level(account.Code) == 1 {
			calcAccounts = append(calcAccounts, account)
			continue
		}
		parent := accountMap[code.Parent(account.Code)]
		parent.Children = append(parent.Children, account)
	}

	var accountDetails = &AccountDetails{}
	// 计算金额
	for _, account := range calcAccounts {
		calc(accountDetails, account)
	}
	accountDetails.Details = calcAccounts
	// 计算总资产、总负债和净资产
	for _, account := range calcAccounts {
		switch account.Type {
		case acc.TypeAsset:
			accountDetails.Total = accountDetails.Total.Add(account.Balance)
		case acc.TypeDebt:
			accountDetails.Debt = accountDetails.Debt.Add(account.Balance)
		}
	}
	accountDetails.NetAmount = accountDetails.Total.Sub(accountDetails.Debt)
	return accountDetails, nil
}

func calc(accountDetails *AccountDetails, account *model.Account) {
	if account.Children == nil {
		return
	}

	for _, child := range account.Children {
		if child.Children != nil {
			calc(accountDetails, child)
		}
		account.Debit = account.Debit.Add(child.Debit)
		account.Credit = account.Credit.Add(child.Credit)
		account.Balance = account.Balance.Add(child.Balance)
	}

}
