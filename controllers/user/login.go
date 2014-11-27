package user

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) Get() {
	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplNames = "user/login.tpl"
}
