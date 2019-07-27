package common

import (
	"client"
	"github.com/juzi5201314/coolq-http-api"
	"time"
)

type SpaceInfo struct {
	Api         cqhttp_go_sdk.API   `json:"api"`
	RedisClient *client.RedisClient `json:"redis_client"`
}

type GroupMessage struct {
	UserId  float64   `json:"user_id"`
	GroupId float64   `json:"group_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
