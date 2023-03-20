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
