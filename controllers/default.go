package controllers

import (
// "fmt"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["title"] = this.Lang("title_login")
	this.TplNames = "index.tpl"
}
