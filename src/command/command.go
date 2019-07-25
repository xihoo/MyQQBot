package command

import (
	"common"
	"github.com/juzi5201314/coolq-http-api/command"
)

func RegisterCommand() {
	command.Register("echo", command.Executant(EchoCommand))
}

func EchoCommand(cmd string, args []string, ctm command.CommandTriggerMan) {
	ctm.Reply(args[0], common.GetSpace().Api)
}
