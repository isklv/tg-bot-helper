package main

import (
	"log"
	"time"

	"github.com/isklv/tg-bot-helper/internal/bot"
	"github.com/isklv/tg-bot-helper/internal/config"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	botInstance := bot.NewBot(cfg.TelegramBotToken)

	err = botInstance.Run()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Duration(1<<63 - 1))
}
