package model

type NovelContent struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	RawContent string `xorm:"comment('未经处理的原始数据') TEXT"`
	Content    string `xorm:"comment('处理好的纯文字内容') TEXT"`
}
