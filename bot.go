package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
	"runtime/debug"
	"strings"
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

	chatBot.Searcher = NewImageSearcher(viper.GetString("telegram.f"))
}

func (chatBot *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	bot := chatBot.TelegramBot
	bot.GetChatMember(tgbotapi.ChatConfigWithUser{})
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		debug.PrintStack()
		return
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		processUpdate(update)

		if strings.HasPrefix(update.Message.Text, "gbr") {

		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.From.UserName, " bilang: ", update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
