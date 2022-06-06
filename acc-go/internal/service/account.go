package service

import (
	"acc/internal/model"
	"acc/internal/pkg/e"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func ListAccount(ownerId, ledgerId int64) ([]*model.Account, error) {

	accounts, err := model.ListAccounts(ledgerId)
	if err != nil {
		logrus.Errorf("Query account failed: %v", err)
		return nil, e.Error
	}

	var accountMap = make(map[int64]*model.Account, len(accounts))

	for _, account := range accounts {
		accountMap[account.ID] = account
	}

	for _, account := range accounts {
		if account.ParentId == 0 {
			continue
		}
		parent := accountMap[account.ParentId]
		parent.Children = append(parent.Children, account)
	}

	var calcAccounts []*model.Account
	for _, account := range accounts {
		if account.ParentId == 0 {
			calcAccounts = append(calcAccounts, account)
		}
	}

	for _, account := range calcAccounts {
		calc(account)
	}

	return calcAccounts, nil
}

func calc(account *model.Account) {
	if account.Children == nil {
		return
	}

	var debit, credit, balance decimal.Decimal
	for _, child := range account.Children {
		if child.Children != nil {
			calc(child)
		}
		debit = debit.Add(child.Debit)
		credit = credit.Add(child.Credit)
		balance = balance.Add(child.Balance)
	}
	account.Debit = debit
	account.Credit = credit
	account.Balance = balance
}
