package dao

import (
	"log"
	"testing"
)

func TestDynuConfig(t *testing.T) {
	username := DynuConfig("username")
	passwd := DynuConfig("password")
	hostname := DynuConfig("hostname")
	log.Printf("username:%s,passwd:%s,hostname:%s", username, passwd, hostname)
}
