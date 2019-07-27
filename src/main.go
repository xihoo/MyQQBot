package main

import (
	"common"
	"flag"
	"github.com/golang/glog"
	"github.com/juzi5201314/coolq-http-api"
	"github.com/juzi5201314/coolq-http-api/cq"
	"github.com/juzi5201314/coolq-http-api/server"
	"strings"
)

func main() {
	configFile := flag.String("conf", "./config.yaml", "config file path")
	flag.Parse()
	common.LoadConfig(*configFile)
	api := cqhttp_go_sdk.Api(common.GetConfig().QQSendAddr+":"+common.GetConfig().QQSendPort, common.GetConfig().QQSendPassword)
	common.InitSpace(api)
	/*_, err := api.SendPrivateMsg(630558072, "hello", false)
	if err != nil {
		fmt.Print(err.Error())
	}*/
	s := server.StartListenServer(8080, common.GetConfig().QQListenAddr)
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
	glog.Info("receive message: " + message + "，type is: " + sub_type + " \n")
	if message[0] == '#' {
		if strings.Contains(message, "dj") {
			return map[string]interface{}{
				"reply":     strings.TrimPrefix(strings.Replace(message, "dj", cq.At("458830973")+" ", -1), "#"),
				"at_sender": false,
			}
		}
		return map[string]interface{}{
			"reply":     strings.TrimPrefix(message, "#"),
			"at_sender": false,
		}
	}
	return nil
}
