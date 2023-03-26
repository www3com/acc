package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func ListAccounts(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.List(ledgerId)
	r.Render(c, accounts, err)
}

func ListIncomes(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.ListIncomes(ledgerId)
	r.Render(c, accounts, err)
}

func ListExpenses(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.ListExpenses(ledgerId)
	r.Render(c, accounts, err)
}

func SaveAccount(c *gin.Context) {
	var account service.Account
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	if account.Id != 0 {
		err := accountService.Update(ledgerId, &account)
		r.Render(c, nil, err)
		return
	}

	err = accountService.Create(ledgerId, &account)
	r.Render(c, nil, err)
}

func DeleteAccount(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var account service.DeleteAccount
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	account.LedgerId = ledgerId
	err = accountService.Delete(&account)
	r.Render(c, nil, err)
}

func UpdateName(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var account service.UpdateName
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	err = accountService.UpdateName(ledgerId, &account)
	r.Render(c, nil, err)
}

func UpdateRemark(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var account service.UpdateRemark
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	err = accountService.UpdateRemark(ledgerId, &account)
	r.Render(c, nil, err)
}

func UpdateBalance(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var account service.UpdateBalance
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	err = accountService.UpdateBalance(ledgerId, &account)
	r.Render(c, nil, err)
}
