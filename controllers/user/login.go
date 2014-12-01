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
	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplNames = "user/login.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	if models.Login(username, password) {
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
}
