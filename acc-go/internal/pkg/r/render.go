package r

import (
	"accounting-service/internal/misc/consts"
	"accounting-service/internal/pkg/e"
	"accounting-service/internal/pkg/i18n"
	"accounting-service/internal/pkg/translator"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

var (
	langKey           = "lang"
	acceptLanguageKey = "Accept-Language"
)

func Render[T any](c *gin.Context, status int, code int, message string, data T) {
	c.JSON(status, &R[T]{
		Code:    code,
		Message: message,
		Data:    data})
}

func RenderOk(c *gin.Context, data interface{}) {
	Render(c, consts.OK, consts.OK, "", data)
}

func RenderCode(c *gin.Context, code int) {
	Render(c, consts.OK, code, getMessage(c, code), "")
}

func RenderError(c *gin.Context, err error) {
	_, ok := err.(*json.UnmarshalTypeError)
	if ok {
		Render(c, consts.INVALID_PARAMS, consts.INVALID_PARAMS, getMessage(c, consts.INVALID_PARAMS), "")
	}

	errs, ok := err.(validator.ValidationErrors)
	if ok {
		Render(c, consts.INVALID_PARAMS, consts.INVALID_PARAMS, translator.Translate(errs), "")
		return
	}

	appErr, ok := err.(*e.AppError)
	if ok {
		RenderCode(c, appErr.Code)
	} else {
		Render(c, consts.INTERNAL_ERROR, consts.INTERNAL_ERROR, getMessage(c, consts.INTERNAL_ERROR), "")
	}
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
