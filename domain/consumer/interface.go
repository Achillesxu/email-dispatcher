package consumer

import "latest/domain/message"

type Repository interface {
	Consumer() (*message.Message, error)
}
