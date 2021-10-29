package repository

import (
	"latest/domain"
	"log"

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

func (m MailRepository) SendMessage(e domain.Message) error {
	y := gomail.NewMessage()
	y.SetHeader("From")
	y.SetHeader("To")
	y.SetHeader("Subject", "")
	y.SetBody("text/html", "")

	if err := m.Dialer.DialAndSend(y); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
