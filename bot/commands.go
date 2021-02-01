package bot

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var CommandsMap = make(map[string]*NormalCommand)

/*type Command interface {
	execute(update tgbotapi.Update) string
}*/

type NormalCommand struct {
	// 名称,匹配命令
	name string
	// 描述,被help命令调用
	desc string
	// 执行的方法
	Call func(update tgbotapi.Update, args ...string) string
}

func Help(m map[string]*NormalCommand) string {
	result := ""
	//var buf bytes.Buffer
	for key, command := range CommandsMap {
		//buf.WriteString(fmt.Sprintf("/%s:%s\n", key, command.desc))
		result = fmt.Sprintf("%s/%s:%s\n", result, key, command.desc)
	}
	return result
}

func init() {
	//CommandsMap
	CommandsMap["json"] = &NormalCommand{
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
	CommandsMap["help"] = &NormalCommand{
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

}
