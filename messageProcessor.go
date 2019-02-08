package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (bot *Bot) ProcessUpdate(update tgbotapi.Update) {
	bot.ProcessMessage(update.Message)
}

func (bot *Bot) ProcessMessage(message *tgbotapi.Message) {
	switch SubstringFirstWord(message.Text) {
	case "gbr":
		if strings.Contains(strings.ToLower(message.Text), "kharriz") {
			bot.SendKont(message)
		} else {
			log.Println("gbr start")
			bot.ProcessGbr(message)
		}
	}
}

func SubstringFirstWord(s string) string {
	ss := strings.SplitN(s, " ", 2)
	if len(ss) < 2 {
		return ""
	}
	return strings.ToLower(ss[0])

}

func (bot *Bot) ProcessGbr(message *tgbotapi.Message) {
	texts := strings.SplitN(message.Text, " ", 2)
	if len(texts) < 2 {
		return
	}
	msg, err := bot.GetImage(texts[1])
	if err == nil {
		msg = msg.SetChatID(message.Chat.ID)
		log.Println("gbr send")
		res, err := bot.TelegramBot.Send(msg.BuildChattable())
		if err != nil {
			log.Println("err: ", err)
		}
		log.Println("res: ", res)
		log.Println("gbr done")
	} else if err == NoSearchResultError {
		res, err := bot.TelegramBot.Send(NewReplyText().SetText("gak ada").SetChatID(message.Chat.ID).BuildChattable())
		if err != nil {
			log.Println("err: ", err)
		}
		log.Println("res: ", res)
	} else {
		log.Println(err)
	}
}

func (bot *Bot) SendKont(message *tgbotapi.Message) {
	bot.TelegramBot.Send(NewReplyText().SetText(message.Chat.UserName + " kontol").SetChatID(message.Chat.ID).BuildChattable())
}

func (bot *Bot) GetImage(query string) (ReplyImage, error) {
	link, err := bot.Searcher.SearchImage(query)
	if err != nil {
		log.Println(err)
		return NewReplyImage(), err
	}
	imageByte, err := DownloadImage(link)
	if err != nil {
		log.Println(err)
		return NewReplyImage(), err
	}
	return NewReplyImage().SetFile(imageByte).SetFileName(query), nil
}

func (bot *Bot) SendMessage() {

}

type Reply interface {
	BuildChattable() tgbotapi.Chattable
}

type ReplyText struct {
	chatID int64
	text   string
}

func (r ReplyText) ChatID() int64 {
	return r.chatID
}

func (r ReplyText) SetChatID(chatID int64) ReplyText {
	r.chatID = chatID
	return r
}

func (r ReplyText) Text() string {
	return r.text
}

func (r ReplyText) SetText(text string) ReplyText {
	r.text = text
	return r
}

func NewReplyText() ReplyText {
	return ReplyText{
		chatID: 0,
		text:   "",
	}
}

func (r ReplyText) BuildChattable() tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(r.chatID, r.text)
	return msg
}

type ReplyImage struct {
	chatID   int64
	file     []byte
	fileName string
}

func NewReplyImage() ReplyImage {
	return ReplyImage{
		chatID:   0,
		file:     nil,
		fileName: "",
	}
}

func (r ReplyImage) BuildChattable() tgbotapi.PhotoConfig {

	b := tgbotapi.FileBytes{Name: "image.jpg", Bytes: r.file}
	msg := tgbotapi.NewPhotoUpload(r.chatID, b)
	msg.Caption = r.fileName
	return msg
}

func (r ReplyImage) FileName() string {
	return r.fileName
}

func (r ReplyImage) SetFileName(fileName string) ReplyImage {
	r.fileName = fileName
	return r
}

func (r ReplyImage) File() []byte {
	return r.file
}

func (r ReplyImage) SetFile(file []byte) ReplyImage {
	r.file = file
	return r
}

func (r ReplyImage) ChatID() int64 {
	return r.chatID
}

func (r ReplyImage) SetChatID(chatID int64) ReplyImage {
	r.chatID = chatID
	return r
}
