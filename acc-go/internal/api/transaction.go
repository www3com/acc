package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if err := c.ShouldBind(&trans); err != nil {
		ret.RenderError(c, err)
		return
	}
	trans.Recorder = 1
	if err := service.CreateTransaction(&trans); err != nil {
		logger.Error("create Transaction error, recorder: {}, details: ", ownerId, err)
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, nil)
}
