package problem

import (
	"github.com/gojudge/ojsite/controllers"
)

type AddProblemController struct {
	controllers.BaseController
}

func (this *AddProblemController) Get() {
	this.Data["title"] = this.Lang("title_add_problem")

	this.TplName = "problem/add.tpl"
}

func (this *AddProblemController) Post() {
	this.Data["title"] = this.Lang("title_add_problem")

	this.TplName = "problem/add.tpl"
}
