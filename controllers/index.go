package controllers

import (
// "fmt"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["title"] = this.Lang("title_index")
	this.TplName = "index.tpl"
}
