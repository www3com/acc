package service

import (
	"acc/internal/consts"
	acc "acc/internal/consts/account"
	"acc/internal/dao"
	"acc/internal/model"
	"acc/internal/pkg/code"
	"github.com/shopspring/decimal"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/log"
	"time"
)

type AccountOverview struct {
	Total     decimal.Decimal `json:"total"`
	Debt      decimal.Decimal `json:"debt"`
	NetAmount decimal.Decimal `json:"netAmount"`
	Details   []*dao.Account  `json:"details"`
}
type AccountService struct{}

// Create 创建账户
func (s *AccountService) Create(ledgerId int64, account *model.AccountBO) error {
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
	dbAccount := dao.Account{
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

// Update 更新账户
func (s *AccountService) Update(ledgerId int64, account *model.UpdateAccountBO) error {
	now := time.Now().UnixMicro()
	cols := make(map[string]interface{})
	cols["update_time"] = now
	switch account.Type {
	case 1:
		cols["name"] = account.Name
	case 2:
		cols["remark"] = account.Remark
	case 3:
	}
	return accountDao.Update(ledgerId, account.Id, cols)
}

func (s *AccountService) Delete(account *model.DeleteAccountBO) error {
	level := code.Level(account.Code)
	if level == 1 {
		return consts.ErrDeleteTopAccount
	}
	count, err := accountDao.ChildCount(account.LedgerId, account.Code, level+1)
	if err != nil {
		return err
	}
	if count > 0 {
		return consts.ErrExistSubAccount
	}

	return accountDao.Delete(account.LedgerId, account.Code)
}

func (s *AccountService) Overview(ledgerId int64) (*AccountOverview, error) {

	accounts, err := accountDao.ListByTypes(ledgerId, acc.TypeAsset, acc.TypeReceivables, acc.TypeDebt)
	if err != nil {
		return nil, e.Wrap("Query account error", err)
	}

	var accountMap = make(map[string]*dao.Account, len(accounts))
	for _, account := range accounts {
		accountMap[account.Code] = account
	}

	// 构造树
	var calcAccounts []*dao.Account
	for _, account := range accounts {
		if code.Level(account.Code) == 1 {
			calcAccounts = append(calcAccounts, account)
			continue
		}
		parent := accountMap[code.Parent(account.Code)]
		parent.Children = append(parent.Children, account)
	}

	var accountDetails = &AccountOverview{}
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

func (s *AccountService) List(ledgerId int64) ([]*dao.Account, error) {
	accounts, err := accountDao.ListByTypes(ledgerId, acc.TypeAsset, acc.TypeReceivables, acc.TypeDebt)
	if err != nil {
		return nil, e.Wrap("list accounts", err)
	}

	return buildTree(accounts), nil
}

func (s *AccountService) ListAll(ledgerId int64) ([]*dao.Account, error) {
	accounts, err := accountDao.ListByTypes(ledgerId)
	if err != nil {
		return nil, e.Wrap("list all accounts", err)
	}

	return buildTree(accounts), nil
}

func calc(accountDetails *AccountOverview, account *dao.Account) {
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

func buildTree(accounts []*dao.Account) []*dao.Account {
	var accountMap = make(map[string]*dao.Account, len(accounts))
	for _, account := range accounts {
		accountMap[account.Code] = account
	}

	// 构造树
	var calcAccounts []*dao.Account
	for _, account := range accounts {
		if code.Level(account.Code) == 1 {
			calcAccounts = append(calcAccounts, account)
			continue
		}
		parent := accountMap[code.Parent(account.Code)]
		parent.Children = append(parent.Children, account)
	}

	return calcAccounts
}
