package mail

import "latest/domain/message"

type Repository interface {
	SendMessage(e message.Message) error
}
