package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetPublicIP() string {
	res, err := http.Get("http://ip.3322.net/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	s, err := ioutil.ReadAll(res.Body)
	return string(s)
}
