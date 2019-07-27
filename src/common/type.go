package common

import (
	"github.com/juzi5201314/coolq-http-api"
	"time"
)

type SpaceInfo struct {
	Api cqhttp_go_sdk.API `json:"api"`
}

type GroupMessage struct {
	UserId  string    `json:"user_id"`
	GroupId string    `json:"group_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
