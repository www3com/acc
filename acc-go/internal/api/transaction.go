package api

import (
	"acc/internal/pkg/context"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var trans service.Transaction
	if err := r.BindAndValid(c, &trans); err != nil {
		r.RenderFail(c, err)
		return
	}
	trans.Recorder = context.GetUserId(c)
	err := service.CreateTransaction(&trans)
	r.Render(c, nil, err)
}
