package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// Controller基类继承封装
type BaseController struct {
	beego.Controller
}

var LANG string

func (this *BaseController) Lang(key string) string {
	lang := this.Ctx.GetCookie("lang")
	if len(lang) == 0 {
		this.Ctx.SetCookie("lang", "zh-CN", 1<<31-1, "/")
		lang = "zh-CN"
	}

	LANG = lang

	return i18n.Tr(lang, key)
}
