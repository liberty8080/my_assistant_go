package boot

import (
	"my_assistant_go/app/model"
	"testing"
)

func TestDbEngine(t *testing.T) {
	var config []model.Config
	d := Engine.Where("type=? and name=?", 2, "username").Find(&config)
	if d != nil {
		t.Error("数据库连接失败")
	}

}

func TestDynuConfig(t *testing.T) {

}
