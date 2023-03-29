package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func ListTransaction(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var transaction service.TransactionBO
	if err := r.BindAndValid(c, &transaction); err != nil {
		r.RenderFail(c, err)
		return
	}

	projects, err := transactionService.List(ledgerId, &transaction)
	r.Render(c, projects, err)
}
