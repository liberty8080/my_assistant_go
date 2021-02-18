package engine

import (
	"log"
	"my_assistant_go/util"
)

func worker(request Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", request.Url)
	content, err := util.GetHtml(request.Url)
	if err != nil {
		log.Printf("请求失败!,Url: %s\n", request.Url)
		return ParseResult{}, err
	}
	return request.ParseFunc(content, request.Params), nil
}
