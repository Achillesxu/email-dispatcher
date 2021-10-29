package usescases

import (
	"fmt"
	"latest/app/dispatcher"
	"latest/dto"
	"log"
)

type Dispatcher struct{}

func NewEmailDispatcher() Usecases {
	return &Dispatcher{}
}

func (d *Dispatcher) Send(dto dto.KafkaResponse, header string) error {

	service := dispatcher.NewMailService()

	mail := service.GetMailInstance()

	if header == "change_email" {

		log.Println(fmt.Sprintf("Recived header: %s", header))

		m := dispatcher.ChangeEmailMessage(dto)

		if err := mail.DialAndSend(m); err != nil {
			log.Println(err)
			return err
		}

		return nil
	}

	if header == "lost_password" {

		log.Println(fmt.Sprintf("Recived header: %s", header))

		m := dispatcher.LostPassMessage(dto)

		if err := mail.DialAndSend(m); err != nil {
			log.Println(err)
			return err
		}

		return nil
	}

	m := dispatcher.CodeMessage(dto)

	if err := mail.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
