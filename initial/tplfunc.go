package initial

/*
import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/gojudge/ojsite/controllers"
	//"github.com/gojudge/ojsite/utils"
	"runtime"
	"strings"
	"time"
)

var (
	AppVer string
)

func i18nGetString(key string) string {
	return i18n.Tr(controllers.LANG, key)
}

func getVer(tag string) string {
	if tag == "app" {
		return AppVer
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

func cut(val string, length int) string {
	if length > len(val) {
		length = len(val)
	}

	return val[0:length]
}

func loadTime(startTime int64) float64 {
	stn := time.Now()
	nowTime := stn.UnixNano()

	return float64(nowTime-startTime) / 1000000000
}

func InitTplFunc() {
	beego.AddFuncMap("i18n", i18nGetString)
	//beego.AddFuncMap("asset", utils.Fis)
	beego.AddFuncMap("ver", getVer)
	beego.AddFuncMap("date", dateOfTime)
	beego.AddFuncMap("cut", cut)
	beego.AddFuncMap("load_time", loadTime)
}
*/
