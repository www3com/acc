package middleware

import (
	"acc/internal/pkg/e"
	t "acc/internal/pkg/jwt"
	"acc/internal/pkg/r"
	"acc/internal/pkg/uid"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader(t.TokenKey)
		if token == "" {
			r.RenderFail(c, e.UnauthorizedError)
			c.Abort()
			return
		}

		parsedToken, err := t.ParseToken(token)
		if err != nil {
			r.RenderFail(c, e.UnauthorizedError)
			c.Abort()
			return
		}

		userId := uid.Uid2Id(parsedToken.Uid)
		c.Header("userId", strconv.FormatInt(userId, 10))

		if parsedToken.IP != c.ClientIP() {
			r.RenderFail(c, e.UnauthorizedError)
			c.Abort()
			return
		}

		c.Next()
	}
}
