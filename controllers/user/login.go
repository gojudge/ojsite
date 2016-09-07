package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/models"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.Forbbiden("not", "guest")

	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplName = "user/login.tpl"
}

func (this *LoginController) Post() {
	this.Forbbiden("not", "guest")

	user := models.User{}

	username := this.GetString("username")
	password := this.GetString("password")

	if result, lev := user.Login(username, password); result {
		user, err := user.GetUser(0, username, "", "")
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"result": false,
				"msg":    "login failed, get user info failed",
				"refer":  nil,
			}
		} else {
			this.SetSession("username", username)
			this.SetSession("userid", fmt.Sprintf("%d", user.Id))
			this.SetSession("level", lev)

			this.Data["json"] = map[string]interface{}{
				"result": true,
				"msg":    "login success",
				"refer":  nil,
			}
		}

	} else {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "login failed",
			"refer":  nil,
		}
	}

	this.ServeJSON()
}
