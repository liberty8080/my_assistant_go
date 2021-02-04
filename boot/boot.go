package boot

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var engine *xorm.Engine

func initDb() {
	var err error
	engine, err = xorm.NewEngine("mysql",
		"aide:jacob_aide@(192.168.98.100:3306)/jacob_aide?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败!")
	}

	// aide:jacob_aide@(192.168.98.100:3306)/jacob_aide?charset=utf8
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
}

func init() {
	initDb()
}
