package api

import (
	"acc/internal/pkg/e"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if code := r.BindAndValid(c, &trans); code != e.OK {
		r.RenderFail(c, code)
		return
	}
	trans.Recorder = 1
	if err := service.CreateTransaction(&trans); err != nil {
		logger.Error("create Transaction e, recorder: {}, details: ", ownerId, err)
		r.RenderFail(c, e.ERROR)
		return
	}
	r.RenderOk(c, nil)
}
