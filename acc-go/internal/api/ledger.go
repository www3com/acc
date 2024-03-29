package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func ListLedger(c *gin.Context) {
	userId, err := context.GetUserId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	token, err := ledgerService.List(userId)
	r.Render(c, token, err)
}
