package main

import (
	"encoding/json"
	"fmt"
	"github.com/juzi5201314/coolq-http-api"
)

func main() {
	fmt.Print("hello world!")
	api := cqhttp_go_sdk.Api("http://localhost:5700", "MAX8char")
	m, err := api.SendPrivateMsg(630558072, "hello", false)
	if err != nil {
		fmt.Print(err.Error())
	}
	mstr, err := json.Marshal(m)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(mstr)
}
