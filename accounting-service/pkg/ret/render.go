package ret

import (
	"accounting-service/pkg/i18n"
	"accounting-service/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

var (
	langKey           = "lang"
	acceptLanguageKey = "Accept-Language"
)

func Render(c *gin.Context, status int, code int, message string, data interface{}) {
	c.JSON(status, &RetData{
		Code:    code,
		Message: message,
		Data:    data})
}

func RenderOk(c *gin.Context, data interface{}) {
	Render(c, http.StatusOK, OK, "", data)
}

func RenderCode(c *gin.Context, code int) {
	Render(c, http.StatusOK, code, getMessage(c, code), nil)
}

func RenderError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		Render(c, http.StatusOK, INVALID_PARAMS, translator.Translate(errs), nil)
	} else {
		Render(c, http.StatusInternalServerError, INTERNAL_ERROR, getMessage(c, INTERNAL_ERROR), nil)
	}
}

func RenderFail(c *gin.Context, status int, code int) {
	Render(c, status, code, "", nil)
}

func getMessage(c *gin.Context, code int) string {
	// 从header中获取指定语言
	lang := c.GetHeader(langKey)
	if lang == "" {
		lang, _ = c.Cookie(langKey)
	}

	// 从参数中获取指定语言
	if lang == "" {
		lang = c.Query(langKey)
	}

	if lang == "" {
		lang = c.GetHeader(acceptLanguageKey)
	}

	msg, _ := i18n.Localize(lang, strconv.Itoa(code), "")
	return msg

}
