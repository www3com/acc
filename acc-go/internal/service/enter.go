package service

import "acc/internal/dao"

var (
	accountDao     = new(dao.AccountDao)
	ledgerDao      = new(dao.LedgerDao)
	memberDao      = new(dao.MemberDao)
	projectDao     = new(dao.ProjectDao)
	supplierDao    = new(dao.SupplierDao)
	userDao        = new(dao.UserDao)
	tplLedgerDao   = new(dao.TplLedgerDao)
	transactionDao = new(dao.TransactionDao)

	ledgerService      = new(LedgerService)
	transactionService = new(TransactionService)
)
