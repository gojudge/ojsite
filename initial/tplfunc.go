package initial

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/utils"
	"runtime"
	"strings"
)

const (
	APP_VER = "0.0.1.1121"
)

func i18nGetString(key string) string {
	return i18n.Tr(controllers.LANG, key)
}

func getVer(tag string) string {
	if tag == "app" {
		return APP_VER
	} else if tag == "golang" {
		return runtime.Version()
	} else {
		return ""
	}

}

func dateOfTime(fullTime string) string {
	date := strings.Split(fullTime, " ")
	return date[0]
}

func InitTplFunc() {
	beego.AddFuncMap("i18n", i18nGetString)
	beego.AddFuncMap("asset", utils.Fis)
	beego.AddFuncMap("ver", getVer)
	beego.AddFuncMap("date", dateOfTime)
}
