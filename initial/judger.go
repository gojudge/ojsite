package initial

import (
	"github.com/astaxie/beego"
	"github.com/duguying/judger/client"
	"github.com/duguying/ojsite/utils"
)

// connect judger and login
func InitJudger() {
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
