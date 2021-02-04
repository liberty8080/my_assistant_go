package boot

import (
	"log"
	"my_assistant_go/app/model"
	"testing"
)

func TestDbEngine(t *testing.T) {
	var config []model.Config
	d := engine.Where("type=? and name=?", 2, "username").Find(&config)
	if d != nil {
		log.Panic("xxx")
	}
	for _, elem := range config {
		log.Println(elem)
	}
}
