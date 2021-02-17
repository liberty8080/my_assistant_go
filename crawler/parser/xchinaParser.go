package parser

import (
	"github.com/antchfx/htmlquery"
	"my_assistant_go/app/model"
	"my_assistant_go/crawler/engine"
	"regexp"
	"strings"
	"time"
)
import _ "github.com/antchfx/htmlquery"

const (
	baseUrl = "https://xchina.co"
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
		//result.Items = append(result.Items, "")
		result.Request = append(result.Request, engine.Request{
			Url:       baseUrl + link,
			ParseFunc: GetFictionInfo,
		})
	}
	return result
}

func GetFictionInfo(html string) engine.ParseResult {
	result := engine.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(html))

	urlNode := htmlquery.FindOne(root, "//meta[@property='og:url']")
	rowId := getRowIdFromUrl(htmlquery.SelectAttr(urlNode, "content"))

	novelNameNode := htmlquery.FindOne(root, "//meta[@property='og:title']")
	novelName := htmlquery.SelectAttr(novelNameNode, "content")

	briefNode := htmlquery.FindOne(root, "//meta[@property='og:description']")
	brief := htmlquery.SelectAttr(briefNode, "content")

	imageNode := htmlquery.FindOne(root, "//meta[@property='og:image']")
	cover := htmlquery.SelectAttr(imageNode, "content")
	// todo:tags,chapters
	novel := model.Novel{
		RawId:      rowId,
		NovelName:  novelName,
		DataSource: "小黄书",
		Brief:      brief,
		Cover:      cover,
		Author:     "",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	result.Items = append(result.Items, novel)
	return result
}

func getRowIdFromUrl(s string) string {
	reg := regexp.MustCompile(`id-(\w+).html`)
	params := reg.FindStringSubmatch(s)
	return params[1]
}
