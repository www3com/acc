package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func ListProjects(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	projects, err := projectService.List(ledgerId)
	r.Render(c, projects, err)
}
