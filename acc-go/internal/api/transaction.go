package api

import (
	"acc/internal/pkg/context"
	"acc/internal/pkg/e"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if err := r.BindAndValid(c, &trans); err != nil {
		r.RenderFail(c, err)
		return
	}
	trans.Recorder = 1
	if err := service.CreateTransaction(&trans); err != nil {
		ownerId := context.GetUserId(c)
		logrus.Error("create Transaction e, recorder: {}, details: ", ownerId, err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, nil)
}
