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

// run before get
func (this *BaseController) Prepare() {
	// get user level
	var lev string

	user := this.GetSession("username")
	if user == nil {
		lev = "guest" // guest, not login
	} else {
		level := this.GetSession("level")

		if level == nil {
			lev = "user"
		} else {
			if tmplev, ok := level.(string); !ok {
				lev = "user"
			} else {
				lev = tmplev
			}
		}
	}

	this.Data["userIs"] = lev
}

// run after finished
func (this *BaseController) Finish() {

}
