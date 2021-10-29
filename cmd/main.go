package main

import (
	"latest/app"
	"latest/config"
)

func main() {

	config.Init()
	config.InitLogger()
	config.Logger().Info("Starting service")

	app.Start()
}
