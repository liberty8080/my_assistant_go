package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"log"
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

func ParseFictionList(html string, params ...interface{}) engine.ParseResult {
	root, _ := htmlquery.Parse(strings.NewReader(html))
	// 小说列表
	list := htmlquery.Find(root, "//div[@class='list']/div[@class='fiction']")
	result := engine.ParseResult{}
	for _, row := range list {
		linkNode := htmlquery.FindOne(row, "./a[contains(@href,'id')]")
		link := htmlquery.SelectAttr(linkNode, "href")
		result.Request = append(result.Request, engine.Request{
			Url:       baseUrl + link,
			ParseFunc: GetFictionInfo,
			//Params: params,
		})
	}
	return result
}

func GetFictionInfo(html string, params ...interface{}) engine.ParseResult {

	result := engine.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(html))

	tagNodes := htmlquery.Find(root, "//div[@class='contentTag']")
	print("tagNodes: ")
	var tags []string
	for _, tag := range tagNodes {
		fmt.Printf("%+v ", htmlquery.InnerText(tag))
		tags = append(tags, htmlquery.InnerText(tag))
	}

	chapterNodes := htmlquery.Find(root, "//div[@class='chapters']/div/a")

	novel := model.Novel{
		RawId:      getRowIdFromUrl(getDataFromMeta(root, "url")),
		NovelName:  getDataFromMeta(root, "title"),
		DataSource: "小黄书",
		Brief:      getDataFromMeta(root, "description"),
		Cover:      getDataFromMeta(root, "image"),
		Author:     "",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	for _, chapterNode := range chapterNodes {
		href := htmlquery.SelectAttr(chapterNode, "href")
		result.Request = append(result.Request, engine.Request{
			Url:       baseUrl + href,
			ParseFunc: GetFictionContent,
			Params:    novel.RawId,
		})
	}
	result.Items = append(result.Items, novel)
	return result
}

func getDataFromMeta(n *html.Node, property string) string {
	node := htmlquery.FindOne(n, fmt.Sprintf("//meta[@property='og:%s']", property))
	return htmlquery.SelectAttr(node, "content")
}

func getRowIdFromUrl(s string) string {
	reg := regexp.MustCompile(`id-(\w+).html`)
	params := reg.FindStringSubmatch(s)
	if len(params) > 0 {
		return params[1]
	} else {
		return ""
	}
}

func GetFictionContent(html string, params ...interface{}) engine.ParseResult {
	result := engine.ParseResult{}
	root, _ := htmlquery.Parse(strings.NewReader(html))
	fictionNode := htmlquery.Find(root, "//div[@class='article large']/div[@class='fiction']/p]")
	rowContent := ""
	content := ""
	for _, node := range fictionNode {
		rowContent += htmlquery.OutputHTML(node, true)
		content += htmlquery.InnerText(node) + "\n"
	}
	chapter := model.NovelChapter{
		NovelId:     params[0].(int),
		ChapterName: getDataFromMeta(root, "title"),
		ContentId:   0,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	novelContent := model.NovelContent{
		RawContent: rowContent,
		Content:    content,
	}
	if params == nil {
		log.Panic("参数传递错误!")
	} else {
		log.Printf("参数:%+v", params)
	}
	result.Items = append(result.Items, novelContent, chapter)
	return result
}
