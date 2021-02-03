package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"my_assistant_go/app/dao"
	"my_assistant_go/util"
	"net/http"
	"testing"
)

func download(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(string(body))
}
func TestNmap(t *testing.T) {
	download("http://ip.3322.net/")
}

func TestPublicIp(t *testing.T) {
	//GetPublicIP()
	println(util.GetPublicIP())
}

func TestHelpCommand(t *testing.T) {

}

func TestExpire(t *testing.T) {
	log.Printf("vmess: %s\n", util.Expire())
}

func TestDb(t *testing.T) {
	//db := g.DB()
	//config := &model.Config{}
	//
	//err := db.Table("config").Struct(config)
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.Println(config)
	config, err := dao.Config.Where("id=?", 3).One()
	if err != nil {
		log.Panic(err)
	}
	log.Println(config)

}
