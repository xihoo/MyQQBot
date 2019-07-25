package client

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
	return nil
}

func (rClient *RedisClient) Get(key string, value interface{}) (interface{}, error) {
	return nil, nil
}

func (rClient *RedisClient) Del(key string, value interface{}) error {
	return nil
}
