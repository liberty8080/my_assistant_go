package bot

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"my_assistant_go/app/dao"
	"my_assistant_go/util"
)

var CommandsMap = make(map[string]*NormalCommand)

type NormalCommand struct {
	// 名称,匹配命令
	name string
	// 描述,被help命令调用
	desc string
	// 执行的方法
	Call func(update tgbotapi.Update, args ...string) (string, error)
}

var Help = &NormalCommand{
	name: "help",
	desc: "帮助信息",
	Call: func(update tgbotapi.Update, args ...string) (string, error) {
		result := ""
		for key, command := range CommandsMap {
			result = fmt.Sprintf("%s/%s: %s\n", result, key, command.desc)
		}
		return result, nil
	},
}

var Json = &NormalCommand{
	name: "json",
	desc: "json数据转换",
	Call: func(update tgbotapi.Update, args ...string) (string, error) {
		jsons, err := json.Marshal(update)
		if err != nil {
			log.Println("json转换失败")
		}
		return string(jsons), err
	},
}

var DDNS = &NormalCommand{
	name: "ddns",
	desc: "同步ddns",
	Call: func(update tgbotapi.Update, args ...string) (string, error) {
		username := dao.DynuConfig("username")
		passwd := dao.DynuConfig("password")
		hostname := dao.DynuConfig("hostname")
		myIp, err := util.GetPublicIP()
		result := util.Get(fmt.Sprintf("https://api.dynu.com/nic/update?hostname=%s&myip=%s&username=%s&password=%s", hostname, myIp, username, passwd))
		return result, err
	},
}

var IP = &NormalCommand{
	name: "ip",
	desc: "获取当前公网ip",
	Call: func(update tgbotapi.Update, args ...string) (string, error) {
		return util.GetPublicIP()
	},
}

var Expire = &NormalCommand{
	name: "expire",
	desc: "查看机场过期时间",
	Call: func(update tgbotapi.Update, args ...string) (string, error) {
		return util.Expire()
	},
}

func init() {
	CommandsMap["json"] = Json
	CommandsMap["help"] = Help
	CommandsMap["ddns"] = DDNS
	CommandsMap["ip"] = IP
	CommandsMap["expire"] = Expire
}
