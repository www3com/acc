package service

import (
	"acc/internal/consts"
	acc "acc/internal/consts/account"
	t "acc/internal/consts/transaction"
	"acc/internal/dao"
	"acc/internal/model"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gorm.io/gorm"
	"time"
)

type TransactionService struct{}

var p = message.NewPrinter(language.English)

func (s *TransactionService) InsertBalance(tx *gorm.DB, ledgerId int64, userId int64, accountId int64, amount decimal.Decimal) error {
	current, err := accountDao.Get(ledgerId, accountId)
	if err != nil {
		return errors.WithMessage(err, "get account when adjusting balance")
	}

	now := time.Now().UnixMicro()
	transDO := dao.Transaction{
		CreateTime: now,
		LedgerId:   ledgerId,
		Amount:     amount.Abs(),
		UserId:     userId,
	}

	switch current.Type {
	case acc.TypeAsset, acc.TypeReceivables:
		// 如果是正数
		if amount.IsPositive() {
			transDO.AccountId = consts.AccountBalanceIncome
			transDO.CpAccountId = accountId
			transDO.Type = t.TypeIncome
		} else {
			transDO.AccountId = consts.AccountBalanceExpenses
			transDO.CpAccountId = accountId
			transDO.Type = t.TypeExpenses
		}
	case acc.TypeDebt, acc.TypeEquity:
		// 如果是正数
		if amount.IsPositive() {
			transDO.AccountId = consts.AccountBalanceExpenses
			transDO.CpAccountId = accountId
			transDO.Type = t.TypeExpenses
		} else {
			transDO.AccountId = consts.AccountBalanceIncome
			transDO.CpAccountId = accountId
			transDO.Type = t.TypeIncome
		}
	}
	if err := transactionDao.Insert(tx, &transDO); err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) Insert(tx *gorm.DB, tran *model.TransactionBO) error {
	now := time.Now().UnixMicro()
	transDO := dao.Transaction{
		CreateTime:  now,
		TradingTime: tran.TradingTime,
		LedgerId:    tran.LedgerId,
		Type:        tran.Type,
		AccountId:   tran.AccountId,
		CpAccountId: tran.CpAccountId,
		Amount:      tran.Amount,
		Remark:      tran.Remark,
		ProjectId:   tran.ProjectId,
		MemberId:    tran.MemberId,
		SupplierId:  tran.SupplierId,
		UserId:      tran.UserId,
	}

	if err := transactionDao.Insert(tx, &transDO); err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) ListTotal(tran *model.TransactionQuery) (*model.TransactionTotalVO, error) {
	totals, err := transactionDao.ListTotal(tran)
	if err != nil {
		return nil, err
	}

	var (
		expense decimal.Decimal
		income  decimal.Decimal
	)
	for _, total := range totals {
		if total.Type == t.TypeExpenses {
			expense = total.Amount
		}
		if total.Type == t.TypeIncome {
			income = total.Amount
		}
	}

	vo := model.TransactionTotalVO{
		Income:  "+" + p.Sprintf("%.2f", income.Round(2).InexactFloat64()),
		Expense: p.Sprintf("%.2f", expense.Round(2).Neg().InexactFloat64()),
		Balance: p.Sprintf("%.2f", income.Sub(expense).Round(2).InexactFloat64()),
	}

	return &vo, nil
}

func (s *TransactionService) List(tran *model.TransactionQuery) ([]*model.TransactionVO, error) {
	dos, err := transactionDao.List(tran)
	if err != nil {
		return nil, err
	}
	accounts, err := accountDao.List(tran.LedgerId)
	if err != nil {
		return nil, err
	}
	projects, err := projectDao.ListAll(tran.LedgerId)
	if err != nil {
		return nil, err
	}
	members, err := memberDao.ListAll(tran.LedgerId)
	if err != nil {
		return nil, err
	}
	suppliers, err := supplierDao.ListAll(tran.LedgerId)
	if err != nil {
		return nil, err
	}
	var vos []*model.TransactionVO
	for _, do := range dos {
		tm := time.UnixMilli(do.TradingTime)
		tran := model.TransactionVO{
			Id:          do.ID,
			TradingTime: tm.Format("2006-01-02"),
			Type:        translateType(do.Type),
			Account:     translateAccount(do.AccountId, accounts),
			CpAccount:   translateAccount(do.CpAccountId, accounts),
			Project:     translateProject(do.ProjectId, projects),
			Member:      translateMember(do.MemberId, members),
			Supplier:    translateSupplier(do.SupplierId, suppliers),
			Amount:      getAmount(do.Type, do.Amount),
			Remark:      do.Remark,
		}
		vos = append(vos, &tran)
	}

	return vos, nil
}

func translateType(typeId int) string {
	return t.Mapper[typeId]
}

func translateAccount(accountId int64, accounts []*dao.Account) string {
	for i := 0; i < len(accounts); i++ {
		if accountId == accounts[i].ID {
			return accounts[i].Name
		}
	}
	return ""
}

func translateProject(projectId int64, projects []*dao.Project) string {
	for i := 0; i < len(projects); i++ {
		if projectId == projects[i].ID {
			return projects[i].Name
		}
	}
	return ""
}

func translateMember(memberId int64, members []*dao.Member) string {
	for i := 0; i < len(members); i++ {
		if memberId == members[i].ID {
			return members[i].Name
		}
	}
	return ""
}

func translateSupplier(supplierId int64, suppliers []*dao.Supplier) string {
	for i := 0; i < len(suppliers); i++ {
		if supplierId == suppliers[i].ID {
			return suppliers[i].Name
		}
	}
	return ""
}

func getAmount(transactionType int, amount decimal.Decimal) string {
	switch transactionType {
	case t.TypeExpenses, t.TypeLend, t.TypeDebtOut, t.TypeTransfer:
		return p.Sprintf("%.2f", amount.Neg().Round(2).InexactFloat64())
	default:
		return p.Sprintf("%.2f", amount.Round(2).InexactFloat64())
	}
}
