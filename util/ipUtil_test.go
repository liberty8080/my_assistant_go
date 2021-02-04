package util

import (
	"log"
	"testing"
)

func TestNmap(t *testing.T) {
	download("http://ip.3322.net/")
}

func TestPublicIp(t *testing.T) {
	//GetPublicIP()
	println(GetPublicIP())
}

func TestHelpCommand(t *testing.T) {

}

func TestExpire(t *testing.T) {
	log.Printf("vmess: %s\n", Expire())
}
