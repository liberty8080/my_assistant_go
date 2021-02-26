package boot

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var Engine *xorm.Engine

func initMysql() (*xorm.Engine, error) {
	return xorm.NewEngine("mysql",
		"aide:jacob_aide@(192.168.98.100:3306)/jacob_aide?charset=utf8")
}

func initSqlite() (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", "")
}

func InitDb() {
	var err error
	Engine, err = initMysql()
	if err != nil {
		log.Fatal("数据库连接失败!")
	}

	// aide:jacob_aide@(192.168.98.100:3306)/jacob_aide?charset=utf8
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
}

func init() {
	InitDb()
}
