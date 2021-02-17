package crawler

import (
	"log"
	"my_assistant_go/crawler/engine"
	"my_assistant_go/crawler/parser"
	"my_assistant_go/util"
	"testing"
)

func TestGetFictionList(t *testing.T) {
	html, err := util.GetHtml("https://xchina.co/fictions/1.html")
	if err != nil {
		log.Panicf("请求失败!")
	}
	result := parser.ParseFictionList(html)
	log.Printf("%v", result)
}

func TestEngine(t *testing.T) {
	e := engine.CrawlerEngine{
		Scheduler:   &engine.SimpleScheduler{},
		WorkerCount: 50,
	}

	e.Run(engine.Request{
		Url:       "https://xchina.co/fictions/1.html",
		ParseFunc: parser.ParseFictionList,
	})
}
