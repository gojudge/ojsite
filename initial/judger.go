package initial

import (
	"github.com/astaxie/beego"
	"github.com/duguying/judger/client"
	"github.com/duguying/ojsite/utils"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
)

// connect judger and login
func InitJudger() {
	host := beego.AppConfig.String("judgerhost")
	port, err := beego.AppConfig.Int("judgerport")
	pass := beego.AppConfig.String("judgerpass")

	if err != nil {
		port = 1004
	}

	err = client.New(host, port)
	if err != nil {
		log.Warnln(err)
		return
	}

	loginInfo := utils.MsgPack(map[string]interface{}{
		"action":   "login",
		"password": pass,
	})

	content, err := client.J.Request(loginInfo)
	if err != nil {
		log.Warnln("[Request]", err)
		return
	}

	data, err := com.JsonDecode(content)
	if err != nil {
		log.Warnln(err)
		return
	}

	json := data.(map[string]interface{})

	if !json["result"].(bool) {
		log.Warnln("Login Judger Failure!")
	}
}
