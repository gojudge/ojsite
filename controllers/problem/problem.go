package problem

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

type ProblemsController struct {
	controllers.BaseController
}

func (this *ProblemsController) Get() {
	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplNames = "user/login.tpl"
}
