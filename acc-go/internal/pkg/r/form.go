package r

import (
	"acc/internal/consts"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/upbos/go-saber/e"
	"strings"
)

var uni *ut.UniversalTranslator

func init() {
	uni = ut.New(en.New(), zh.New())
}

func BindAndValid(c *gin.Context, p interface{}) error {
	err := c.ShouldBind(p)
	if err == nil {
		return nil
	}

	locale := c.GetHeader(consts.Locale)
	trans, _ := uni.GetTranslator(locale)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		switch locale {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			_ = zhTranslations.RegisterDefaultTranslations(v, trans)
		}
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var arr []string
	for _, value := range errs.Translate(trans) {
		arr = append(arr, value)
	}

	return e.WithMessage(e.InvalidParams, strings.Join(arr, ","))
}
