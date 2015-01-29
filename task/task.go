package task

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/duguying/judger/client"
	"github.com/duguying/ojsite/utils"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"time"
)

func CheckJudger() error {
	log.Blueln(fmt.Sprintf("[%s]", com.SubString(time.Now().String(), 0, 19)), "check judger")

	ping := utils.MsgPack(map[string]interface{}{
		"action": "ping",
	})
	_, err := client.J.Request(ping)
	if err != nil {
		reconnect()
	}

	return nil
}

func reconnect() {
	host := beego.AppConfig.String("judgerhost")
	port, err := beego.AppConfig.Int("judgerport")
	pass := beego.AppConfig.String("judgerpass")

	if err != nil {
		port = 1004
	}

	client.New(host, port)
	loginInfo := utils.MsgPack(map[string]interface{}{
		"action":   "login",
		"password": pass,
	})

	client.J.Request(loginInfo)
}
