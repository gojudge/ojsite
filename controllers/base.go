package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/duguying/ojsite/models"
	"github.com/gogather/com"
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

func (this *BaseController) Forbbiden(condition string) {
	if "logout" == condition {
		if this.Data["userIs"] == "guest" { //logout
			this.Redirect("/", 302)
		}
	} else if "login" == condition {
		if this.Data["userIs"] != "guest" { // login
			this.Redirect("/", 302)
		}
	} else if "student" == condition {
		if this.Data["userIs"] == "student" { //logout
			this.Redirect("/", 302)
		}
	} else if "teacher" == condition {
		if this.Data["userIs"] == "teacher" { //logout
			this.Redirect("/", 302)
		}
	} else if "admin" == condition {
		if this.Data["userIs"] == "admin" { //logout
			this.Redirect("/", 302)
		}
	}
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

		username := user.(string)
		u, err := models.GetUser(0, username, "", "")
		if err != nil {
			this.Data["nickname"] = ""
			this.Data["email_md5"] = ""
		} else {
			this.Data["nickname"] = u.Nickname
			this.Data["email_md5"] = com.Md5(u.Email)
		}

	}

	this.Data["userIs"] = lev
}

// run after finished
func (this *BaseController) Finish() {

}
