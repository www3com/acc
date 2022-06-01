package middleware

import (
	"acc/internal/pkg/e"
	t "acc/internal/pkg/jwt"
	"acc/internal/pkg/r"
	"acc/internal/pkg/uid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader(t.ACC_TOKEN_KEY)
		if token == "" {
			r.Render(c, http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		parsedToken, err := t.ParseToken(token)
		if err != nil {
			r.Render(c, http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		userId := uid.Uid2Id(parsedToken.Uid)
		c.Header("userId", strconv.FormatInt(userId, 10))

		if parsedToken.IP != c.ClientIP() {
			r.Render(c, http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
