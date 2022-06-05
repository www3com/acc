package api

import (
	"acc/internal/pkg/context"
	"acc/internal/pkg/e"
	r "acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// CreateLedger 创建用户账本
func CreateLedger(c *gin.Context) {
	if err := service.CreateLedger(1, "标准账本", context.GetUserId(c)); err != nil {
		logrus.Errorf("createLedger API, error: %v", err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, nil)
}

// UpdateLedger 更新用户账本的名称
func UpdateLedger(c *gin.Context) {
	if ledger, ok := valid(c); ok {
		if err := service.UpdateLedger(ledger.LedgerId, ledger.Name); err != nil {
			logrus.Error("Update Ledger name e, ledger id: {}, details: ", ledger.LedgerId, err)
			r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		r.RenderOk(c, nil)
	}
}

// DeleteLedger 删除用户账本
func DeleteLedger(c *gin.Context) {
	if ledger, ok := valid(c); ok {
		if err := service.DeleteLedger(ledger.LedgerId); err != nil {
			logrus.Error(err)
			r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
			return
		}
		r.RenderOk(c, nil)
	}
}

func ListLedger(c *gin.Context) {
	userId := context.GetUserId(c)
	ledgers, err := service.ListLedger(userId)
	if err != nil {
		logrus.Error("Query user account e, user id: {}, details: ", userId, err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, ledgers)
}

func valid(c *gin.Context) (*service.Ledger, bool) {
	var ledger service.Ledger

	if err := r.BindAndValid(c, &ledger); err != nil {
		r.RenderFail(c, err)
		return nil, false
	}
	return &ledger, true
}
