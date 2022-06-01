package context

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const HEADER_USER_ID = "userId"

func GetUserId(c *gin.Context) int64 {
	userId := c.GetHeader(HEADER_USER_ID)
	id, _ := strconv.ParseInt(userId, 10, 64)
	return id
}
