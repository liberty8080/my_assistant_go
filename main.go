package main

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"my_assistant_go/app/dao"
	"my_assistant_go/bot"
	"os"
	"strconv"
	"strings"
)

var BOT *tgBotApi.BotAPI

func handleUpdate(update tgBotApi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	if strings.HasPrefix(update.Message.Text, "/") {
		for name, handler := range bot.CommandsMap {
			if name == update.Message.Text[1:] {
				//参数形式待定
				log.Print("执行" + name)
				msg := tgBotApi.NewMessage(update.Message.Chat.ID, handler.Call(update))
				_, _ = BOT.Send(msg)
			}
		}
	}
}

func InitBot() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("请设置环境变量BOT_TOKEN")
	}
	botAPI, err := tgBotApi.NewBotAPI(token)
	if err != nil {
		log.Println(os.Getenv("BOT_TOKEN"))
		log.Panic(err)
	}
	BOT = botAPI
	botAPI.Debug = true

	log.Printf("成功登录BOT: %s ", botAPI.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60
	adminID, _ := strconv.Atoi(dao.BotConfig("admin_id"))
	updates := botAPI.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.From.ID != adminID {
			errMsg := "用户:" + update.Message.From.UserName +
				" : " + update.Message.Text
			log.Printf(errMsg)
			_, _ = botAPI.Send(tgBotApi.NewMessage(int64(adminID), errMsg))
			continue
		}
		handleUpdate(update)
	}
}

func main() {
	InitBot()
	log.Print("starting")
}
