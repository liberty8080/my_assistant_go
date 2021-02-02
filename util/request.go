package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) string {
	res, err := http.Get("https://subscribe.zealingcloud.info/link/Zxv1RYVrJ7MXhSFS?sub=3")
	if err != nil {
		log.Panic("订阅内容请求失败!")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic("http response body read failed")
	}
	return string(body)
}
