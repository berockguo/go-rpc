package rpc

import (
	"strings"

	"gitlab.ucloudadmin.com/ucre/rpc/common/util"

	"github.com/berockguo/goini"
)

type Config struct {
	ServiceName string
}

func NewConfig(name string) *Config {
	return &Config{ServiceName: name}
}

func (o *Config) ServerAddress() (string, error) {
	conf := goini.SetConfig("etc/" + o.ServiceName + ".conf")
	ip := conf.GetValue("Server", "ip")
	port := conf.GetValue("Server", "port")
	if strings.Compare(port, "no value") == 0 {
		return port, &util.Error{"empty value"}
	}
	s := ip + ":" + port
	return s, nil
}
func (o *Config) ServerPort() (string, error) {
	conf := goini.SetConfig("etc/" + o.ServiceName + ".conf")
	port := conf.GetValue("Server", "port")
	if strings.Compare(port, "no value") == 0 {
		return port, &util.Error{"empty value"}
	}
	s := ":" + port
	return s, nil
}
