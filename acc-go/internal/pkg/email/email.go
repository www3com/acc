package email

import (
	"acc/internal/pkg/setting"
	"bytes"
	"gopkg.in/gomail.v2"
	"text/template"
)

func SendTemplateMail[T](path string, mailAddr string, subject string, data T) {
	t, err := template.ParseFiles(path)
	var tmplBytes bytes.Buffer
	err = t.Execute(&tmplBytes, data)
	if err != nil {
		panic(err)
	}
	SendMail(mailAddr, subject, tmplBytes.String())
}

func SendMail(mailAddr string, subject string, body string) error {
	m := gomail.NewMessage()

	from := "Acc Center" + "<" + setting.EmailSetting.UserName + ">"
	m.SetHeader("From", from)
	m.SetHeader("To", mailAddr)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(setting.EmailSetting.Host, setting.EmailSetting.Port, setting.EmailSetting.UserName, setting.EmailSetting.Password)
	err := d.DialAndSend(m)
	return err
}
