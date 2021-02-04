package model

type ConfigType struct {
	Id       int    `xorm:"INT"`
	TypeName string `xorm:"VARCHAR(255)"`
	Comment  string `xorm:"VARCHAR(255)"`
}
