package repository

import (
	"context"
	"encoding/json"
	"latest/config"
	domain "latest/domain/message"
	"latest/dto"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.GetConfig().Brokers},
			GroupID: config.GetConfig().Topic,
			Topic:   config.GetConfig().Topic}),
	}
}

func (k *Consumer) Consumer() (*domain.Message, error) {

	ctx := context.Background()

	dto := dto.KafkaResponse{}

	message, err := k.reader.ReadMessage(ctx)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(message.Value, &dto)

	if err != nil {
		return nil, err
	}

	dmn := domain.NewMessage(config.GetConfig().MailUser, dto.To, dto.Subject, dto.Template)

	return &dmn, nil
}
