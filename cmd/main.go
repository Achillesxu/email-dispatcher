package main

import (
	"latest/config"
	"latest/usecases/kafka"
	"log"
)

func main() {

	config.Init()

	log.Println("Started...")
	log.Println(config.GetConfig().Topic)

	kafka := kafka.NewKafkaConsumer()

	kafka.Consumer()
}
