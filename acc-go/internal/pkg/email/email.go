package email

import (
	"acc/internal/pkg/setting"
	"bytes"
	"gopkg.in/gomail.v2"
	"text/template"
)

func SendTemplateMail[T any](path string, mailAddr string, subject string, data T) error {
	t, err := template.ParseFiles(path)
	var tmplBytes bytes.Buffer
	err = t.Execute(&tmplBytes, data)
	if err != nil {
		panic(err)
	}
	return SendMail(mailAddr, subject, tmplBytes.String())
}

func SendMail(mailAddr string, subject string, body string) error {
	m := gomail.NewMessage()

	from := "Acc Center" + "<" + setting.ConfigSetting.Email.User + ">"
	m.SetHeader("From", from)
	m.SetHeader("To", mailAddr)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(setting.ConfigSetting.Email.Host, setting.ConfigSetting.Email.Port, setting.ConfigSetting.Email.User, setting.ConfigSetting.Email.Password)
	err := d.DialAndSend(m)
	return err
}
