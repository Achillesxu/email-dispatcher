package main

import (
	"latest/app"
	"latest/config"
	"log"
)

func main() {

	config.Init()

	log.Println("Started...")
	log.Println(config.GetConfig().Topic)

	app.Start()
}
