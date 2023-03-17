package service

import "acc/internal/model"

var (
	accountDao   = new(model.AccountDao)
	ledgerDao    = new(model.LedgerDao)
	memberDao    = new(model.MemberDao)
	projectDao   = new(model.MemberDao)
	supplierDao  = new(model.SupplierDao)
	userDao      = new(model.UserDao)
	tplLedgerDao = new(model.TplLedgerDao)

	ledgerService = new(LedgerService)
)
