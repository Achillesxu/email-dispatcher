package dispatcher

import (
	"latest/domain/consumer"
	"latest/domain/mail"
)

type Dispatcher struct {
	mail     mail.Repository
	consumer consumer.Repository
}

func NewEmailDispatcher(mail mail.Repository, consumer consumer.Repository) Usecases {
	return &Dispatcher{
		mail:     mail,
		consumer: consumer,
	}
}

func (d *Dispatcher) Dispatch() {
	for {

		msg, err := d.consumer.Consumer()

		if err != nil {
			continue
		}

		err = d.mail.SendMessage(*msg)

		if err != nil {
			continue
		}

	}

}
