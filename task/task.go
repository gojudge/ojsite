package task

import (
	"github.com/astaxie/beego"
	"github.com/duguying/judger/client"
	"github.com/duguying/ojsite/utils"
	"github.com/gogather/com/log"
)

func CheckJudger() error {
	ping := utils.MsgPack(map[string]interface{}{
		"action": "ping",
	})
	_, err := client.J.Request(ping)
	if err != nil {
		log.Warnln("reconnecting")
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
