package usescases

import "latest/dto"

type Usecases interface {
	Send(dto dto.KafkaResponse, header string) error
}
