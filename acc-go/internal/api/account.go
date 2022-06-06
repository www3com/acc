package api

import (
	"acc/internal/pkg/context"
	"acc/internal/pkg/r"
	"acc/internal/service"
	"github.com/gin-gonic/gin"
)

func ListAccount(c *gin.Context) {
	ownerId := context.GetUserId(c)
	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		r.RenderFail(c, err)
	} else {
		r.RenderOk(c, accounts)
	}
}
