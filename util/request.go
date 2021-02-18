package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Panicf("get请求失败!url:%s", url)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic("http response body read failed")
	}
	return string(body)
}

func GetHtml(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	var (
		err     error
		resp    *http.Response
		retries = 3
	)
	for retries > 0 {
		client := &http.Client{Timeout: time.Second * 5}
		resp, err = client.Do(req)
		if err != nil {
			log.Printf("请求失败,一秒后重新请求,错误信息:%v", err)
			time.Sleep(time.Second)
			retries -= 1
		} else {
			break
		}
	}

	if resp != nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil && data == nil {
			log.Panicln(err)
		}
		return fmt.Sprintf("%s", data), nil
	}
	return "", err

}
