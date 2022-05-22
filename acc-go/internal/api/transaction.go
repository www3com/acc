package api

import (
	"accounting-service/internal/misc/consts"
	"accounting-service/internal/pkg/logger"
	"accounting-service/internal/pkg/r"
	"accounting-service/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if err := c.ShouldBind(&trans); err != nil {
		r.RenderError(c, err)
		return
	}
	trans.Recorder = 1
	if err := service.CreateTransaction(&trans); err != nil {
		logger.Error("create Transaction e, recorder: {}, details: ", ownerId, err)
		r.RenderCode(c, consts.INTERNAL_ERROR)
		return
	}
	r.RenderOk(c, nil)
}
