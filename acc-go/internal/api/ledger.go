package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func ListLedger(c *gin.Context) {
	userId := context.GetUserId(c)
	token, err := ledgerService.List(userId)
	r.Render(c, token, err)
}
