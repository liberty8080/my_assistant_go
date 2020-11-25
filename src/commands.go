package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var Commands map[string]func(update tgbotapi.Update) string

func toJson(update tgbotapi.Update) string {
	jsons, err := json.Marshal(update)
	if err != nil {
		log.Println("json转换失败")
	}
	return string(jsons)
}

func init() {
	Commands["json"] = toJson

}
