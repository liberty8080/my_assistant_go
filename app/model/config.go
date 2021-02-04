package model

type Config struct {
	Id    int    `xorm:"INT"`
	Name  string `xorm:"VARCHAR(255)"`
	Value string `xorm:"VARCHAR(255)"`
	Type  int    `xorm:"INT"`
}
