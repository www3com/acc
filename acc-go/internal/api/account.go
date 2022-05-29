package api

import (
	"acc/internal/pkg/e"
	"acc/internal/pkg/logger"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func ListAccount(c *gin.Context) {
	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		logger.Error("Query user account e, user id: {}, details: ", ownerId, err)
		r.RenderFail(c, e.ERROR)
		return
	}
	r.RenderOk(c, accounts)
}
