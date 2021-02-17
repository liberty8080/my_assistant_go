package parser

import (
	"github.com/antchfx/htmlquery"
	"my_assistant_go/crawler/engine"
	"strings"
)
import _ "github.com/antchfx/htmlquery"

const (
	novelStartPage = "https://xchina.co/fictions/1.html"
	baseUrl        = "https://xchina.co"
)

func ParseFictionList(html string) engine.ParseResult {
	root, _ := htmlquery.Parse(strings.NewReader(html))
	// 小说列表
	list := htmlquery.Find(root, "//div[@class='list']/div[@class='fiction']")
	result := engine.ParseResult{}
	for _, row := range list {
		//log.Println(htmlquery.OutputHTML(row, true))
		linkNode := htmlquery.FindOne(row, "./a[contains(@href,'id')]")
		link := htmlquery.SelectAttr(linkNode, "href")
		//log.Println(htmlquery.SelectAttr(link, "href"))
		result.Items = append(result.Items, "小说id:"+htmlquery.InnerText(linkNode))
		result.Request = append(result.Request, engine.Request{
			Url:       baseUrl + link,
			ParseFunc: GetFictionInfo,
		})
	}
	return result
}

func GetFictionInfo(html string) engine.ParseResult {
	result := engine.ParseResult{}

	result.Items = append(result.Items, html)
	return result
}
