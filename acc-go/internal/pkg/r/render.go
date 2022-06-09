package r

import (
	"acc/internal/pkg/e"
	"acc/internal/pkg/i18n"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	langKey           = "lang"
	acceptLanguageKey = "Accept-Language"
)

func Render(c *gin.Context, data interface{}, err error) {
	if err == nil {
		RenderOk(c, data)
	} else {
		RenderFail(c, err)
	}
}

func RenderOk(c *gin.Context, data interface{}) {
	render(c, http.StatusOK, e.OK, data)
}

func RenderFail(c *gin.Context, err error) {
	v, ok := err.(*e.WrapError)
	if !ok {
		logrus.Errorf("%v", err)
		render(c, http.StatusInternalServerError, e.ERROR, nil)
	}

	switch v.Code {
	case e.ERROR:
		render(c, http.StatusInternalServerError, v.Code, nil)
	case e.InvalidParams:
		render(c, http.StatusBadRequest, v.Code, nil)
	case e.UNAUTHORIZED:
		render(c, http.StatusUnauthorized, v.Code, nil)
	case e.FORBIDDEN:
		render(c, http.StatusForbidden, v.Code, nil)
	case e.NotFound:
		render(c, http.StatusNotFound, v.Code, nil)
	default:
		render(c, http.StatusOK, v.Code, nil)
	}
}

func render(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, &R{
		Code:    code,
		Message: getMessage(c, code),
		Data:    data})
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

	return i18n.GetMessage(lang, code)
}
