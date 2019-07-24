package main

import (
	"fmt"
	"github.com/juzi5201314/coolq-http-api"
	"github.com/juzi5201314/coolq-http-api/server"
)

func main() {
	api := cqhttp_go_sdk.Api("http://localhost:5700", "MAX8char")
	_, err := api.SendPrivateMsg(630558072, "hello", false)
	if err != nil {
		fmt.Print(err.Error())
	}
	s := server.StartListenServer(5700, "/")
	s.ListenPrivateMessage(server.PrivateMessageListener(pm))
	s.Listen()
}

func pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {
	println("receive message: " + message + "，type is" + sub_type)
	return map[string]interface{}{
		"reply": "auto reply",
	}
}
