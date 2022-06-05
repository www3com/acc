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

func Render(c *gin.Context, status int, code int, data interface{}) {
	c.JSON(status, &R{
		Code:    code,
		Message: getMessage(c, code),
		Data:    data})
	return
}

func RenderOk(c *gin.Context, data interface{}) {
	Render(c, http.StatusOK, e.OK, data)
}

func RenderFail(c *gin.Context, err error) {
	v, ok := err.(*e.Error)
	if !ok {
		logrus.Errorf("%+v", err)
		Render(c, http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	switch v.Code {
	case e.ERROR:
		logrus.Errorf("%+v", v.Cause)
		Render(c, http.StatusInternalServerError, v.Code, nil)
	case e.INVALID_PARAMS:
		Render(c, http.StatusBadRequest, v.Code, nil)
	case e.UNAUTHORIZED:
		Render(c, http.StatusUnauthorized, v.Code, nil)
	case e.FORBIDDEN:
		Render(c, http.StatusForbidden, v.Code, nil)
	case e.NOT_FOUND:
		Render(c, http.StatusNotFound, v.Code, nil)
	default:
		Render(c, http.StatusOK, v.Code, nil)
	}
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
