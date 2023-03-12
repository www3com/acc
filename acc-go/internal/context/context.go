package context

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	headerUserId = "userId"
	ledgerId     = "ledgerId"
)

func GetUserId(c *gin.Context) int64 {
	userId := c.GetHeader(headerUserId)
	id, _ := strconv.ParseInt(userId, 10, 64)
	return id
}

func GetLedgerId(c *gin.Context) int64 {
	ledgerId := c.Query(ledgerId)
	id, _ := strconv.ParseInt(ledgerId, 10, 64)
	return id
}
