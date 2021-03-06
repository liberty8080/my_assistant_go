package model

import (
	"time"
)

type NovelChapter struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	NovelId     int       `xorm:"index INT"`
	ChapterName string    `xorm:"VARCHAR(255)"`
	ContentId   int       `xorm:"index INT"`
	CreateTime  time.Time `xorm:"DATETIME"`
	UpdateTime  time.Time `xorm:"DATETIME"`
}
