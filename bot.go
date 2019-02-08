package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
)

type Bot struct {
	TelegramBot *tgbotapi.BotAPI
	Searcher    *ImageSearcher
}

func NewBot(bot *tgbotapi.BotAPI, searcher *ImageSearcher) *Bot {
	return &Bot{
		TelegramBot: bot,
		Searcher:    searcher,
	}
}

func (chatBot *Bot) InitConfig() {
	telegramToken := viper.GetString("telegram.token")
	if telegramToken == "" {
		log.Fatalf("Telegram token are not set in config.toml")
	}
	var err error = nil
	chatBot.TelegramBot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatalf("Telegram token are not set in config.toml")
	}

	//chatBot.Searcher = NewImageSearcher(viper.GetString("telegram.f"))
}

func (bot *Bot) Start() {

	tBot := bot.TelegramBot
	log.Printf("Authorized on account %s", tBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	tBot.GetChatMember(tgbotapi.ChatConfigWithUser{})
	updates, err := tBot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		bot.ProcessUpdate(update)
	}
}
