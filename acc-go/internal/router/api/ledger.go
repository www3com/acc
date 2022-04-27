package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

// CreateLedger 创建用户账本
func CreateLedger(c *gin.Context) {
	var ledger service.Ledger
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

// UpdateLedger 更新用户账本的名称
func UpdateLedger(c *gin.Context) {
	var ledger service.Ledger
	if err := c.ShouldBindQuery(&ledger); err != nil {
		ret.RenderError(c, err)
		return
	}

	if err := service.UpdateLedger(ledger.LedgerId, ledger.Name); err != nil {
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, nil)
}

// DeleteLedger 删除用户账本
func DeleteLedger(c *gin.Context) {
	var ledger service.Ledger
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
