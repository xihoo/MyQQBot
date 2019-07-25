package common

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RedisAddr      string `yaml:"redisAddr"`
	RedisPort      string `yaml:"redisPort"`
	RedisPassword  string `yaml:"redisPassword"`
	QQListenAddr   string `yaml:"qqListenAddr"`
	QQListenPort   string `yaml:"qqListenPort"`
	QQSendAddr     string `yaml:"qqSendAddr"`
	QQSendPort     string `yaml:"qqSendPort"`
	QQSendPassword string `yaml:"qqSendPassword"`
}

var Configure Config

func GetConfig() *Config {
	return &Configure
}

func LoadConfig(filename string) {
	configByte, err := ioutil.ReadFile(filename)
	if err != nil {
		glog.Fatalf("[Robot] Read Config File Error: %s", err.Error())
	}
	err = yaml.Unmarshal(configByte, &Configure)
	if err != nil {
		glog.Fatalf("[Robot] Unmarshal Config Error : %s", err.Error())
	}
}
