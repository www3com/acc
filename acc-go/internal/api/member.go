package api

import (
	"acc/internal/context"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
)

func ListMembers(c *gin.Context) {
	ledgerId, err := context.GetLedgerId(c)
	if err != nil {
		r.RenderFail(c, err)
		return
	}
	members, err := memberService.List(ledgerId)
	r.Render(c, members, err)
}
