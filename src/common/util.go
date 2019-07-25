package common

import "github.com/juzi5201314/coolq-http-api"

var Space SpaceInfo

func GetSpace() *SpaceInfo {
	return &Space
}

func InitSpace(api cqhttp_go_sdk.API) {
	Space.Api = api
}
