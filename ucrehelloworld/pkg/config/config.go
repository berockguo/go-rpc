package config

import (
	"strings"

	"github.com/berockguo/goini"
)

func ServerAddress() (string, bool) {
	conf := goini.SetConfig("etc/ucrehelloworld.conf")
	ip := conf.GetValue("Server", "ip")
	port := conf.GetValue("Server", "port")
	if strings.Compare(port, "no value") == 0 {
		return port, false
	}
	s := ip + ":" + port
	return s, true
}
func ServerPort() (string, bool) {
	conf := goini.SetConfig("etc/ucrehelloworld.conf")
	port := conf.GetValue("Server", "port")
	if strings.Compare(port, "no value") == 0 {
		return port, false
	}
	s := ":" + port
	return s, true
}
