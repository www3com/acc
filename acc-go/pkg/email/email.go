package email

import (
	"accounting-service/pkg/setting"
	"gopkg.in/gomail.v2"
)

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
