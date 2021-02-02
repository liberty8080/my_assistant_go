package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"my_assistant_go/bot"
	"os"
	"strings"
)

var BOT *tgbotapi.BotAPI

func handleUpdate(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	if strings.HasPrefix(update.Message.Text, "/") {
		for name, handler := range bot.CommandsMap {
			if name == update.Message.Text[1:] {
				//参数形式待定
				log.Print("执行" + name)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, handler.Call(update))
				_, _ = BOT.Send(msg)
			}
		}
	}
}

func init() {
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	BOT = botAPI
	botAPI.Debug = false

	log.Printf("成功登录BOT: %s", botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		handleUpdate(update)
	}
}

func main() {
	log.Print("starting")
}
