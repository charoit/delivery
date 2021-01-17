package main

import (
	"delivery/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
