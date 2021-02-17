package parser

import (
	"log"
	"my_assistant_go/util"
	"testing"
)

func TestGetFictionInfo(t *testing.T) {
	html, _ := util.GetHtml("https://xchina.co/fiction/id-602ce920a648f.html")
	result := GetFictionInfo(html)
	log.Printf("result:%v", result)
}
