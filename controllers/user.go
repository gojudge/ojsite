package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {

	this.Data["title"] = this.Lang("title_login")
	this.Data["github_client_id"] = beego.AppConfig.String("github_client_id")
	this.TplNames = "user/login.tpl"
}
