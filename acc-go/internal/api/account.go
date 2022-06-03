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

func ListAccount(c *gin.Context) {
	ownerId := context.GetUserId(c)
	accounts, err := service.ListAccount(ownerId, 7)
	if err != nil {
		logrus.Errorf("Query account api, error: %v", err)
		r.Render(c, http.StatusInternalServerError, e.ERROR, nil)
	} else {
		r.RenderOk(c, accounts)
	}
}
