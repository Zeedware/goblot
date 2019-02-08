package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("./config.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	telegramToken := viper.GetString("telegram.token")
	if telegramToken == "" {
		log.Fatalf("Telegram token are not set in config.toml")
	}
	log.Println(telegramToken)
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}
	imageSearchConfig := NewImageSearchConfig().
		SetUrl(viper.GetString("contexttualwebsearch.url")).
		SetKey(viper.GetString("contexttualwebsearch.key"))
	imageSearcher := NewImageSearcher()
	imageSearcher.setConfig(imageSearchConfig)

	bot.Debug = true

	chatBot := NewBot(bot, imageSearcher)
	chatBot.Start()

}

func SearchImage() {

}
