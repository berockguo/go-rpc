package rpc

import (
	"strings"

	"gitlab.ucloudadmin.com/ucre/rpc/pkg/common/util"

	"github.com/berockguo/goini"
)

type Config struct {
	ServiceName string
}

func NewConfig(name string) *Config {
	return &Config{ServiceName: name}
}

func (o *Config) ServerAddress() (string, error) {
	conf := goini.SetConfig("etc/route.conf")
	ip := conf.GetValue(o.ServiceName, "ip")
	port := conf.GetValue(o.ServiceName, "port")
	if strings.Compare(port, "no value") == 0 {
		return port, &util.Error{"empty value"}
	}
	s := ip + ":" + port
	return s, nil
}
func (o *Config) ServerPort() (string, error) {
	conf := goini.SetConfig("etc/route.conf")
	port := conf.GetValue(o.ServiceName, "port")
	if strings.Compare(port, "no value") == 0 {
		return port, &util.Error{"empty value"}
	}
	s := ":" + port
	return s, nil
}
