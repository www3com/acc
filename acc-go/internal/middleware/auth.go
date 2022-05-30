package middleware

import (
	"acc/internal/pkg/e"
	t "acc/internal/pkg/jwt"
	"acc/internal/pkg/r"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("ACC-TOKEN")
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

		println(c.ClientIP())
		if parsedToken.IP != c.ClientIP() {
			r.Render(c, http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
