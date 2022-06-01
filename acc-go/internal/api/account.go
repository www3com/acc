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

func ListAccount(c *gin.Context) {
	ownerId := context.GetUserId(c)
	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		logger.Error("Query user account e, user id: {}, details: ", ownerId, err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	r.RenderOk(c, accounts)
}
