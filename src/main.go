package main

import (
	"fmt"
	"github.com/juzi5201314/coolq-http-api"
	"github.com/juzi5201314/coolq-http-api/server"
	"strings"
)

func main() {
	api := cqhttp_go_sdk.Api("http://localhost:5700", "MAX8char")
	_, err := api.SendPrivateMsg(630558072, "hello", false)
	if err != nil {
		fmt.Print(err.Error())
	}
	s := server.StartListenServer(8080, "/msg")
	s.ListenPrivateMessage(server.PrivateMessageListener(pm))
	s.ListenGroupMessage(server.GroupMessageListener(gm))
	s.Listen()
}

func pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {
	//fmt.Print("receive message: " + message + "，type is" + sub_type)
	if strings.Contains(message, "#") {
		return map[string]interface{}{
			"reply": strings.TrimPrefix(message, "#"),
		}
	}
	return nil
}

func gm(sub_type string, message_id float64, group_id float64, user_id float64, anonymous string, message string, font float64) map[string]interface{} {
	fmt.Print("receive message: " + message + "，type is: " + sub_type + " \n")
	if strings.Contains(message, "#") {
		return map[string]interface{}{
			"reply": message,
		}
	}
	return nil
}
