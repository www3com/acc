package api

import (
	"acc/internal/context"
	"acc/internal/model"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func Overview(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.Overview(ledgerId)
	r.Render(c, accounts, err)
}

func ListAccounts(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.List(ledgerId)
	r.Render(c, accounts, err)
}

func ListAllAccounts(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	accounts, err := accountService.ListAll(ledgerId)
	r.Render(c, accounts, err)
}

func CreateAccount(c *gin.Context) {
	var account model.AccountBO
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	err = accountService.Create(ledgerId, &account)
	r.Render(c, nil, err)
}

func UpdateAccount(c *gin.Context) {

	var account model.UpdateAccountBO
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	userId, err := context.GetUserId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	err = accountService.Update(ledgerId, userId, &account)
	r.Render(c, nil, err)
}

func DeleteAccount(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var account model.DeleteAccountBO
	if err := r.BindAndValid(c, &account); err != nil {
		r.RenderFail(c, err)
		return
	}

	account.LedgerId = ledgerId
	err = accountService.Delete(&account)
	r.Render(c, nil, err)
}
