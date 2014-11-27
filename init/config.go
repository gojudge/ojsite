package init

import (
	"github.com/astaxie/beego"
	"github.com/gogather/com"
)

func InitConfig() {
	if com.FileExist("custom/app.conf") {
		beego.AppConfigPath = "custom/app.conf"
	} else {
		beego.AppConfigPath = "conf/app.conf"
	}
}
