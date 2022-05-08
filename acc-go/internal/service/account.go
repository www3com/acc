package service

import (
	"accounting-service/internal/model"
	"github.com/shopspring/decimal"
)

func ListAccount(ownerId, ledgerId int64) ([]*model.Account, error) {

	accounts, err := model.ListAccount(ledgerId)
	if err != nil {
		return nil, err
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
		//if parent.Children == nil {
		//	parent.Children = []*model.Account{}
		//}
		parent.Children = append(parent.Children, account)
	}

	var calcedAccounts []*model.Account
	for _, account := range accounts {
		if account.ParentId == 0 {
			calcedAccounts = append(calcedAccounts, account)
		}
	}

	for _, account := range calcedAccounts {
		calc(account)
	}

	return calcedAccounts, nil
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
