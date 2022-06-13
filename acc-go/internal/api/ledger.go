package api

import (
	"acc/internal/pkg/context"
	r "acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

// CreateLedger 创建用户账本
func CreateLedger(c *gin.Context) {
	err := service.CreateLedger(1, "标准账本", "", context.GetUserId(c))
	r.Render(c, nil, err)
}

// UpdateLedger 更新用户账本的名称
func UpdateLedger(c *gin.Context) {
	var ledger service.Ledger

	if err := r.BindAndValid(c, &ledger); err != nil {
		r.RenderFail(c, err)
		return
	}

	err := service.UpdateLedger(ledger.LedgerId, ledger.Name)
	r.Render(c, nil, err)
}

// DeleteLedger 删除用户账本
func DeleteLedger(c *gin.Context) {

	var ledger service.Ledger
	if err := r.BindAndValid(c, &ledger); err != nil {
		r.RenderFail(c, err)
		return
	}

	err := service.DeleteLedger(ledger.LedgerId)
	r.Render(c, nil, err)

}

func ListLedger(c *gin.Context) {
	userId := context.GetUserId(c)
	ledgers, err := service.ListLedgers(userId)
	r.Render(c, ledgers, err)
}
