package util

import (
	"log"
	"my_assistant_go/app/dao"
	"testing"
)

func TestNmap(t *testing.T) {
	download("http://ip.3322.net/")
}

func TestPublicIp(t *testing.T) {
	//GetPublicIP()
	println(GetPublicIP())
}

func TestHelpCommand(t *testing.T) {

}

func TestExpire(t *testing.T) {
	log.Printf("vmess: %s\n", Expire())
}

func TestDb(t *testing.T) {
	//db := g.DB()
	//config := &model.Config{}
	//
	//err := db.Table("config").Struct(config)
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.Println(config)
	config, err := dao.Config.Where("id=?", 3).One()
	if err != nil {
		log.Panic(err)
	}
	log.Println(config)

}
