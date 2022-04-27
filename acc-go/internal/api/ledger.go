package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

const ownerId = 1

// CreateLedger 创建用户账本
func CreateLedger(c *gin.Context) {
	if ledger, ok := valid(c); ok {
		if err := service.CreateLedger(ledger.LedgerId, ownerId); err != nil {
			logger.Error(err)
			ret.RenderCode(c, ret.INTERNAL_ERROR)
			return
		}
		ret.RenderOk(c, nil)
	}
}

// UpdateLedger 更新用户账本的名称
func UpdateLedger(c *gin.Context) {
	if ledger, ok := valid(c); ok {
		if err := service.UpdateLedger(ledger.LedgerId, ledger.Name); err != nil {
			logger.Error("Update Ledger name error, ledger id: {}, details: ", ledger.LedgerId, err)
			ret.RenderCode(c, ret.INTERNAL_ERROR)
			return
		}
		ret.RenderOk(c, nil)
	}
}

// DeleteLedger 删除用户账本
func DeleteLedger(c *gin.Context) {
	if ledger, ok := valid(c); ok {
		if err := service.DeleteLedger(ledger.LedgerId); err != nil {
			logger.Error(err)
			ret.RenderCode(c, ret.INTERNAL_ERROR)
			return
		}
		ret.RenderOk(c, nil)
	}
}

func ListLedger(c *gin.Context) {
	ledgers, err := service.ListLedger(ownerId)
	if err != nil {
		logger.Error("Query user account error, user id: {}, details: ", ownerId, err)
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, ledgers)
}

func valid(c *gin.Context) (*service.Ledger, bool) {
	var ledger service.Ledger
	if err := c.ShouldBindQuery(&ledger); err != nil {
		logger.Error(err)
		ret.RenderError(c, err)
		return nil, false
	}
	return &ledger, true
}
