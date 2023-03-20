package r

import (
	"acc/internal/consts"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"github.com/upbos/go-saber/e"
	"github.com/upbos/go-saber/log"
	"golang.org/x/text/language"
	"net/http"
	"strconv"
)

var mapper = map[int]int{
	e.ERROR:        http.StatusInternalServerError,
	e.Unauthorized: http.StatusUnauthorized,
	e.Forbidden:    http.StatusForbidden,
	e.NotFound:     http.StatusNotFound,
}

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./assets/acc.en.toml")
	bundle.MustLoadMessageFile("./assets/acc.zh.toml")
}

func RenderOk(c *gin.Context, data interface{}) {
	Render(c, data, nil)
}

func RenderFail(c *gin.Context, err error) {
	Render(c, nil, err)
}

func Render(c *gin.Context, data interface{}, err error) {
	if err == nil {
		render(c, http.StatusOK, e.OK, nil, data)
		return
	}

	v, ok := e.UnWrap(err)
	if !ok {
		log.Error(err, "")
		render(c, http.StatusInternalServerError, e.ERROR, nil, data)
		return
	}

	if v.Code == e.InvalidParams {
		render(c, http.StatusBadRequest, v.Code, map[string]interface{}{"err": v.Message}, nil)
		return
	}

	if status := mapper[v.Code]; status == 0 {
		render(c, http.StatusOK, v.Code, v.Data, data)
	} else {
		render(c, status, v.Code, v.Data, data)
	}
}

func render(c *gin.Context, status int, code int, messageData interface{}, data interface{}) {
	c.JSON(status, &R{
		Code:    code,
		Message: getMessage(c, code, messageData),
		Data:    data})
}

func getMessage(c *gin.Context, code int, messageData interface{}) string {
	if code == e.OK {
		return ""
	}
	lang := c.GetHeader(consts.Locale)
	accept := c.GetHeader(consts.AcceptLanguage)
	localizer := i18n.NewLocalizer(bundle, lang, accept)

	msg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: strconv.Itoa(code),
		},
		TemplateData: messageData,
	})

	return msg
}
