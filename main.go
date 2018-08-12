package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	viper.SetConfigFile("./config.toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	telegramToken:=viper.GetString("telegram.token")
	if telegramToken == ""{
		log.Fatalf("Telegram token are not set in config.toml")
	}
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	bot.GetChatMember(tgbotapi.ChatConfigWithUser{})
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if(strings.HasPrefix(update.Message.Text, "gbr")){

		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.From.UserName, " bilang: ", update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}

}

func SearchImage(){

}