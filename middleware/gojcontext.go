package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/beego/i18n"
	"github.com/labstack/echo"
)

type OJContext struct {
	echo.Context
}

func (c OJContext) Trans(key string) string {
	log.Info(key)
	var langValue string
	lang, _ := c.Cookie("lang")
	if lang == nil {
		langValue = "zh-CN"
		lang1 := new(echo.Cookie)
		lang1.SetName(langValue)
		lang1.SetValue("zh-CN")
		lang1.SetPath("/")
		c.SetCookie(lang1)
	} else {
		langValue = lang.Value()
	}

	return i18n.Tr(langValue, key)
}

func OJHandle(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &OJContext{c}
		res := make(map[string]interface{})
		res["i18n"] = cc.Trans
		cc.Set("res", res)
		return h(cc)
	}
}
