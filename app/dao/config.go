package dao

import (
	"log"
	"my_assistant_go/app/model"
	"my_assistant_go/boot"
)

func DynuConfig(name string) string {
	config := new(model.Config)
	_, err := boot.Engine.Where("type=2 and name=?", name).Get(config)
	if err != nil {
		log.Panic("配置查询失败!", err)
	}

	return config.Value
}
