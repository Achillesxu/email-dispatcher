package main

import (
	"latest/config"
	"log"
)

func main() {

	config.Init()

	log.Println("Started...")
	log.Println(config.GetConfig().Topic)

}
