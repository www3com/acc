package api

import (
	"acc/internal/pkg/context"
	"acc/internal/pkg/e"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if code := r.BindAndValid(c, &trans); code != e.OK {
		r.RenderFail(c, code)
		return
	}
	trans.Recorder = 1
	if err := service.CreateTransaction(&trans); err != nil {
		ownerId := context.GetUserId(c)
		logger.Error("create Transaction e, recorder: {}, details: ", ownerId, err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, nil)
}
