package controllers

import (
// "github.com/astaxie/beego"
)

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Data["title"] = "404"
	this.TplNames = "error/404.tpl"
}

func (this *ErrorController) Error501() {
	this.Data["title"] = "501"
	this.TplNames = "error/501.tpl"
}
