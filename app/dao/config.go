package dao

import (
	"log"
	"my_assistant_go/app/model"
	"my_assistant_go/boot"
)

func DynuConfig(name string) string {
	return GetConfig(name, 2)
}

func BotConfig(name string) string {
	return GetConfig(name, 1)
}

func GetConfig(name string, typeNum int) string {
	config := new(model.Config)
	_, err := boot.Engine.Where("type=? and name=?", typeNum, name).Get(config)
	if err != nil {
		log.Panic("配置查询失败!", err)
	}

	return config.Value
}
