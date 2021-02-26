package parser

import (
	"log"
	"my_assistant_go/app/dao"
	"my_assistant_go/app/model"
	"my_assistant_go/util"
	"testing"
)

func TestGetFictionInfo(t *testing.T) {
	html, err := util.GetHtml("https://xchina.co/fiction/id-602ce920a648f.html")
	if err != nil {
		t.Error("爬取失败")
		log.Panic(err)
	}
	result := GetFictionInfo(html)
	dao.AddOrUpdateNovel(result.Items[0].(*model.Novel))
	log.Printf("result:%+v", result)
}

func TestGetFictionContent(t *testing.T) {
	html, err := util.GetHtml("https://xchina.co/fiction/id-602e2caf08235/1.html")
	if err != nil {
		t.Error("爬取失败")
		log.Panic(err)
	}
	result := GetFictionContent(html, 1)
	log.Printf("result:%+v", result)
}
