package context

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const headerUserId = "userId"

func GetUserId(c *gin.Context) int64 {
	userId := c.GetHeader(headerUserId)
	id, _ := strconv.ParseInt(userId, 10, 64)
	return id
}
