package api

import (
	"accounting-service/internal/service"
	"accounting-service/pkg/logger"
	"accounting-service/pkg/ret"
	"github.com/gin-gonic/gin"
)

func ListAccount(c *gin.Context) {
	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		logger.Error("Query user account error, user id: {}, details: ", ownerId, err)
		ret.RenderCode(c, ret.INTERNAL_ERROR)
		return
	}
	ret.RenderOk(c, accounts)
}
