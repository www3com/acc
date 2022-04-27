package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

type Ledger struct {
	LedgerId int64 `form:"ledgerId" binding:"required"`
}

func CreateLedger(c *gin.Context) {
	var ledger Ledger
	if err := c.ShouldBindQuery(&ledger); err != nil {
		ret.RenderError(c, err)
		return
	}

	if err := service.CreateLedger(ledger.LedgerId, 1); err != nil {
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, nil)
}

func UpdateLedger(c *gin.Context) {

}

func DeleteLedger(c *gin.Context) {
	var ledger Ledger
	if err := c.ShouldBindQuery(&ledger); err != nil {
		ret.RenderError(c, err)
		return
	}

	if err := service.DeleteLedger(ledger.LedgerId); err != nil {
		ret.RenderCode(c, ret.INTERNAL_ERROR)
	}
	ret.RenderOk(c, nil)
}

func ListLedger(c *gin.Context) {

}
