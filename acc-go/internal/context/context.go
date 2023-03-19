package context

import (
	"acc/internal/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	headerUserId = "userId"
	ledgerId     = "ledgerId"
)

var ledgerDao = new(model.LedgerDao)

func GetUserId(c *gin.Context) (int64, error) {
	userId := c.GetHeader(headerUserId)
	id, err := strconv.ParseInt(userId, 10, 64)
	return id, err
}

func GetLedgerId(c *gin.Context) (int64, error) {
	userId, err := GetUserId(c)
	if err != nil {
		return 0, err
	}
	ledger, err := ledgerDao.Default(userId)
	if err != nil {
		return 0, err
	}
	return ledger.ID, nil
}
