package api

import (
	"acc/internal/context"
	"acc/internal/model"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
	"github.com/upbos/go-saber/db"
)

func ListTransaction(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var transaction model.TransactionQuery
	if err := r.BindAndValid(c, &transaction); err != nil {
		r.RenderFail(c, err)
		return
	}
	transaction.LedgerId = ledgerId
	trans, err := transactionService.List(&transaction)
	r.Render(c, trans, err)
}

func ListTotalTransaction(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}

	var transaction model.TransactionQuery
	if err := r.BindAndValid(c, &transaction); err != nil {
		r.RenderFail(c, err)
		return
	}
	transaction.LedgerId = ledgerId
	trans, err := transactionService.ListTotal(&transaction)
	r.Render(c, trans, err)
}

func SaveTransaction(c *gin.Context) {
	var tran model.TransactionBO
	if err := r.BindAndValid(c, &tran); err != nil {
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
	tran.LedgerId = ledgerId
	tran.UserId = userId
	err = transactionService.Insert(db.DB, &tran)
	r.Render(c, nil, err)
}
