package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
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
