package model

import (
	"time"
)

type Novel struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	RawId      string    `xorm:"comment('来源网站的小说id') VARCHAR(255)"`
	NovelName  string    `xorm:"comment('小说名称') VARCHAR(255)"`
	DataSource string    `xorm:"comment('来源网站') VARCHAR(255)"`
	Brief      string    `xorm:"comment('简介') VARCHAR(2000)"`
	Cover      string    `xorm:"comment('封面') VARCHAR(255)"`
	Author     string    `xorm:"comment('作者') VARCHAR(20)"`
	RawUrl     string    `xorm:"VARCHAR(255)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateTime time.Time `xorm:"DATETIME"`
}
