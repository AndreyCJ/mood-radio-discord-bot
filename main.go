package main

import (
	"fmt"

	"github.com/AndreyCJ/mood-radio-discord-bot/bot"
	"github.com/AndreyCJ/mood-radio-discord-bot/config"
)

func main() {
	configErr := config.ReadConfig()

	if configErr != nil {
		fmt.Println(configErr.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
}
