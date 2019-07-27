package preprocess

import (
	"common"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/juzi5201314/coolq-http-api/cq"
	"strconv"
	"strings"
	"time"
)

func Gm(sub_type string, message_id float64, group_id float64, user_id float64, anonymous string, message string, font float64) map[string]interface{} {
	glog.Info("receive message: " + message + "，type is: " + sub_type + " \n")
	msg := &common.GroupMessage{
		UserId:  user_id,
		GroupId: group_id,
		Message: message,
		Time:    time.Now(),
	}
	msgByte, _ := json.Marshal(msg)
	key := strconv.FormatFloat(group_id, 'f', -1, 64) + strconv.Itoa(time.Now().Nanosecond())
	err := common.GetSpace().RedisClient.Set(key, string(msgByte))
	if err != nil {
		glog.Errorf("set Message error : %s", err.Error())
	}
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

func Pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {
	//fmt.Print("receive message: " + message + "，type is" + sub_type)
	if strings.Contains(message, "#") {
		return map[string]interface{}{
			"reply": strings.TrimPrefix(message, "#"),
		}
	}
	return nil
}
