package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

var BOT *tgbotapi.BotAPI

func handleUpdate(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	if strings.HasPrefix(update.Message.Text, "/") {
		for name, handler := range CommandsMap {
			if name == update.Message.Text[1:] {
				//参数形式待定
				log.Print("执行" + name[1:])
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, handler.call(update))

				msg.ReplyToMessageID = update.Message.MessageID

				_, _ = BOT.Send(msg)
			}
		}
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("1230304283:AAH4Gu7pw4IPjDXn0_mJf9x9aiHOmF-sj2Q")
	if err != nil {
		log.Panic(err)
	}
	BOT = bot
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		handleUpdate(update)

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
