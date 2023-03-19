package api

import "acc/internal/service"

var (
	userService    = new(service.UserService)
	ledgerService  = new(service.LedgerService)
	accountService = new(service.AccountService)
)
