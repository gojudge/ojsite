package problem

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
)

type ProblemsController struct {
	controllers.BaseController
}

func (this *ProblemsController) Get() {
	pro := &models.Problem{}
	problems, hasNext, _, _ := pro.ListProblem(1, 10, "public")

	this.Data["problems"] = problems
	this.Data["title"] = this.Lang("title_problems")
	this.Data["has_next"] = hasNext
	this.TplNames = "problem/list.tpl"
}
