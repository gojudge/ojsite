package user

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.Forbbiden("login")

	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplNames = "user/login.tpl"
}

func (this *LoginController) Post() {
	this.Forbbiden("login")

	username := this.GetString("username")
	password := this.GetString("password")

	if result, lev := models.Login(username, password); result {
		this.SetSession("username", username)
		this.SetSession("level", lev)

		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "login success",
			"refer":  nil,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "login failed",
			"refer":  nil,
		}
	}

	this.ServeJson()
}
