package controllers

import (
// "github.com/astaxie/beego"
)

type RegistorController struct {
	BaseController
}

func (this *RegistorController) Get() {
	this.Data["title"] = this.Lang("title_registor")
	this.TplNames = "user/registor.tpl"
}
