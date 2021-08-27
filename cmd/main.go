package main

import (
	"latest/app/kafka"
	"latest/config"
	"log"
)

func main() {

	config.Init()

	log.Println("Started...")
	log.Println(config.GetConfig().Topic)

	kafka := kafka.NewKafkaConsumer()

	kafka.Consumer()
}
