package usescases

import (
	"latest/dto"
)

type Dispatcher struct{}

func NewEmailDispatcher() Usecases {
	return &Dispatcher{}
}

func (d *Dispatcher) Send(dto dto.KafkaResponse, header string) error {

	return nil
}
