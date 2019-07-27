package client

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	RedisAddr string `json:"redis_addr"`
	Port      string `json:"port"`
	Password  string `json:"password"`
}

func NewRedisClient(ip string, port string, password string) *RedisClient {
	return &RedisClient{
		RedisAddr: ip,
		Port:      port,
		Password:  password,
	}
}

func (rClient *RedisClient) Set(key string, value interface{}) error {
	conn, err := redis.Dial("tcp", rClient.RedisAddr+":"+rClient.Port)
	if err != nil {
		return err
	}
	var valueStr string
	switch arg := value.(type) {
	case string:
		valueStr = arg
	default:
		return errors.New("value type not supported")
	}
	reply, err := conn.Do("set", key, valueStr)
	if err != nil {
		return err
	}
	if reply != "OK" {
		return errors.New("reply is not OK")
	}
	return nil
}

func (rClient *RedisClient) Get(key string, value interface{}) (interface{}, error) {
	return nil, nil
}

func (rClient *RedisClient) Del(key string, value interface{}) error {
	return nil
}
