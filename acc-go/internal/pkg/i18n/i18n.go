package i18n

import (
	"golang.org/x/text/language"
)

var supported = []language.Tag{
	language.AmericanEnglish,   // 英语
	language.SimplifiedChinese, // 简体中文
	language.Chinese,           // 中文
}

var matcher = language.NewMatcher(supported)

func GetMessage(lang string, key int) string {
	tags, _, _ := language.ParseAcceptLanguage(lang)
	tag, _, _ := matcher.Match(tags...)

	switch tag.Parent() {
	case language.SimplifiedChinese:
		return zh[key]
	case language.Chinese:
		return zh[key]
	case language.English:
		return en[key]
	case language.AmericanEnglish:
		return en[key]
	}
	return zh[key]
}
