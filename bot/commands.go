package bot

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var CommandsMap = make(map[string]*NormalCommand)

type NormalCommand struct {
	// 名称,匹配命令
	name string
	// 描述,被help命令调用
	desc string
	// 执行的方法
	Call func(update tgbotapi.Update, args ...string) string
}

var Help = &NormalCommand{
	name: "help",
	desc: "帮助信息",
	Call: func(update tgbotapi.Update, args ...string) string {
		result := ""
		//var buf bytes.Buffer
		for key, command := range CommandsMap {
			//buf.WriteString(fmt.Sprintf("/%s:%s\n", key, command.desc))
			result = fmt.Sprintf("%s/%s:%s\n", result, key, command.desc)
		}
		return result
	},
}

var Json = &NormalCommand{
	name: "json",
	desc: "json数据转换",
	Call: func(update tgbotapi.Update, args ...string) string {
		jsons, err := json.Marshal(update)
		if err != nil {
			log.Println("json转换失败")
		}
		return string(jsons)
	},
}

func init() {
	//CommandsMap
	CommandsMap["json"] = Json
	CommandsMap["help"] = Help

}
