package initial

import (
	log "github.com/Sirupsen/logrus"
	"github.com/beego/i18n"
	"github.com/gojudge/ojsite/global"
	"strings"
)

// 初始化多国语言模块
func InitLang() {
	languages, _ := global.Config.GetValue("config", "langs")
	if len(languages) <= 0 {
		languages = "zh-CN"
	}
	langs := strings.Split(languages, "|")

	for _, lang := range langs {
		log.Info("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/locale/"+"locale_"+lang+".ini"); err != nil {
			log.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
