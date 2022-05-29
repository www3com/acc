package r

import (
	"acc/internal/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	langKey           = "lang"
	acceptLanguageKey = "Accept-Language"
)

func Render(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, &Msg{
		Code:    code,
		Message: getMessage(c, code),
		Data:    data})
	return
}

func RenderOk(c *gin.Context, data interface{}) {
	Render(c, http.StatusOK, e.OK, data)
}

func RenderFail(c *gin.Context, code int) {
	Render(c, http.StatusOK, code, nil)
}

func getMessage(c *gin.Context, code int) string {
	// 从header中获取指定语言
	lang := c.GetHeader(langKey)
	if lang == "" {
		lang, _ = c.Cookie(langKey)
	}

	if lang == "" {
		lang = c.GetHeader(acceptLanguageKey)
	}

	if lang == "" {
		lang = "zh"
	}

	return e.GetMessage(lang, code)
}
