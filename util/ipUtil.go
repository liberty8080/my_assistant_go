package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// 获取当前公网ip
func GetPublicIP() string {
	res, err := http.Get("http://ip.42.pl/raw")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	s, err := ioutil.ReadAll(res.Body)
	return string(s)
}

type VmessObj struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	TLS        string `json:"tls"`
	VerifyCert bool   `json:"verify_cert"`
	Add        string `json:"add"`
	Port       int    `json:"port"`
	Aid        int    `json:"aid"`
	Net        string `json:"net"`
	HeaderType string `json:"headerType"`
	V          string `json:"v"`
	Type       string `json:"type"`
	Ps         string `json:"ps"`
	Remark     string `json:"remark"`
	ID         string `json:"id"`
	Class      int    `json:"class"`
}

func Expire() string {
	res := Get("https://subscribe.zealingcloud.info/link/Zxv1RYVrJ7MXhSFS?sub=3")
	encodedLinks, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		log.Panic("解码失败!")
	}
	var result = ""
	for _, encoded := range strings.Split(string(encodedLinks), "\n") {
		if len(encoded) != 0 {
			link, err := base64.StdEncoding.DecodeString(strings.Replace(encoded, "vmess://", "", 1))
			if err != nil {
				log.Panic("vmess解码失败! vmess:" + string(link))
			}
			//log.Println(string(link))
			v := VmessObj{}
			_ = json.Unmarshal(link, &v)
			if strings.Contains(v.Remark, "剩余流量") || strings.Contains(v.Remark, "过期时间") {
				//log.Println(v.Remark)
				result += v.Remark + "\n"
			}
		}
	}
	return result
}
