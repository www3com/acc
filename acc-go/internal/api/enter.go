package api

import "acc/internal/service"

var (
	userService     = new(service.UserService)
	ledgerService   = new(service.LedgerService)
	accountService  = new(service.AccountService)
	memberService   = new(service.MemberService)
	projectService  = new(service.ProjectService)
	supplierService = new(service.SupplierService)
)
