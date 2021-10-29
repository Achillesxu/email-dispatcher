package repository

import (
	"latest/config"
	"latest/domain"

	"gopkg.in/gomail.v2"
)

type MailRepository struct {
	*gomail.Dialer
}

func NewMailRepository(mail *domain.Mail) MailRepository {
	return MailRepository{
		Dialer: gomail.NewDialer(mail.Smtp(), int(mail.Port()), mail.User(), mail.Password()),
	}
}

func (m MailRepository) SendMessage() *gomail.Message {
	y := gomail.NewMessage()
	y.SetHeader("From", config.GetConfig().MailUser)
	y.SetHeader("To", d.Email)
	y.SetHeader("Subject", "Greetings from us - Go service")
	y.SetBody("text/html", str)

	return y
}
