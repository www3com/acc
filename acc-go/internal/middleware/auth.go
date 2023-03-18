package middleware

import (
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/jwt"
	"github.com/upbos/go-saber/uid"
	"strconv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader(jwt.TokenKey)
		if token == "" {
			r.RenderFail(c, e.ErrUnauthorized)
			c.Abort()
			return
		}

		parsedToken, err := jwt.ParseToken(token)
		if err != nil {
			r.RenderFail(c, e.ErrUnauthorized)
			c.Abort()
			return
		}

		userId := uid.Id(parsedToken.Uid)
		c.Request.Header.Set("userId", strconv.FormatInt(userId, 10))

		if parsedToken.IP != c.ClientIP() {
			r.RenderFail(c, e.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
