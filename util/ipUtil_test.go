package util

import (
	"log"
	"testing"
)

func TestNmap(t *testing.T) {
	download("http://ip.3322.net/")
}

func TestPublicIp(t *testing.T) {
	println(GetPublicIP())
}

func TestHelpCommand(t *testing.T) {

}

func TestExpire(t *testing.T) {
	result, err := Expire()
	if err != nil {
		log.Printf("Log if")
	}
	log.Printf("vmess: %s\n", result)
}
