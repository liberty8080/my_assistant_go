package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func InitBot() {
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Println(os.Getenv("BOT_TOKEN"))
		log.Panic(err)
	}
	BOT = botAPI
	botAPI.Debug = false

	log.Printf("成功登录BOT: %s ", botAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.From.ID != 551322172 {
			_, _ = botAPI.Send(tgbotapi.NewMessage(551322172, "用户:"+update.Message.From.UserName+
				" : "+update.Message.Text))
			continue
		}
		handleUpdate(update)
	}
}

func main() {
	//boot.InitDb()
	//boot.InitBot()
	InitBot()
	log.Print("starting")
}
