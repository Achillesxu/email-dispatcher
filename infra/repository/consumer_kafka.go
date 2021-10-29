package repository

import (
	"context"
	"encoding/json"
	"latest/config"
	"latest/dto"
	"latest/services"
	"log"

	"github.com/segmentio/kafka-go"
)

type kafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer() *kafkaConsumer {
	return &kafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.GetConfig().Brokers},
			GroupID: config.GetConfig().Topic,
			Topic:   config.GetConfig().Topic}),
	}
}

func (k *kafkaConsumer) Consumer() (string, error) {

	ctx := context.Background()

	dto := dto.KafkaResponse{}

	for {

		message, err := k.reader.ReadMessage(ctx)

		if err != nil {
			continue
		}

		header := message.Headers[0].Key

		err = json.Unmarshal(message.Value, &dto)

		if err != nil {
			continue
		}

		service := services.Dispatcher{}

		err = service.Send(dto, header)

		if err != nil {
			log.Println(err)
			continue

		}
	}
}
