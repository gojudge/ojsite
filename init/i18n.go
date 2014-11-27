package init

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
)

// 初始化多国语言模块
func InitLang() {
	languages := beego.AppConfig.String("langs")
	langs := strings.Split(languages, "|")

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/locale/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
