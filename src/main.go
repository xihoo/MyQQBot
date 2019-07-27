package main

import (
	"client"
	"common"
	"flag"
	"github.com/juzi5201314/coolq-http-api"
	"github.com/juzi5201314/coolq-http-api/server"
	"preprocess"
	"strconv"
)

func main() {
	configFile := flag.String("conf", "./config.yaml", "config file path")
	flag.Parse()
	common.LoadConfig(*configFile)
	api := cqhttp_go_sdk.Api(common.GetConfig().QQSendAddr+":"+common.GetConfig().QQSendPort, common.GetConfig().QQSendPassword)
	redisClient := client.NewRedisClient(common.GetConfig().RedisAddr, common.GetConfig().RedisPort, common.GetConfig().RedisPassword)
	common.InitSpace(api, redisClient)
	/*_, err := api.SendPrivateMsg(630558072, "hello", false)
	if err != nil {
		fmt.Print(err.Error())
	}*/
	port, _ := strconv.Atoi(common.GetConfig().QQListenPort)
	s := server.StartListenServer(port, common.GetConfig().QQListenAddr)
	//s.ListenPrivateMessage(server.PrivateMessageListener(common.Pm))
	s.ListenGroupMessage(server.GroupMessageListener(preprocess.Gm))
	s.Listen()
}
