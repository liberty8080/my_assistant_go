package util

import (
	"io/ioutil"
	"log"
	"net/http"
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
