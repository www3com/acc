package translator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func Setup() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zhTranslations.RegisterDefaultTranslations(validate, trans)
}

//Translate 翻译错误信息
func Translate(errs validator.ValidationErrors) string {
	var result []string
	for _, fieldErr := range errs {
		result = append(result, fieldErr.Translate(trans))
	}

	return strings.Join(result, "; ")
}
