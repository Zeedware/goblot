package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func processUpdate(update tgbotapi.Update) {
	processMessage(update.Message)
}

func processMessage(message *tgbotapi.Message) {
	if strings.HasPrefix(message.Text, "gbr") {

	}
}

func getImages() {

}

type reply struct {
}
