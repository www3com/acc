package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

const defaultLang = "zh"

func Setup() {
	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("assets/locale/zh.toml")

}

// Localize 按数组顺序翻译指定语言
func Localize(lang string, key string, data interface{}) (string, error) {
	if lang == "" {
		lang = defaultLang
	}
	localizer := i18n.NewLocalizer(bundle, lang)
	return localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key,
		},
		TemplateData: data,
	})
}
