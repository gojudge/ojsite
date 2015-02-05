package task

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/judger"
	"github.com/gogather/com/log"
)

func CheckJudger() error {
	jdg := judger.Get("default")
	err := jdg.Ping()
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

	jdg := judger.Get("default")
	jdg.Start(host, port, pass)

}
