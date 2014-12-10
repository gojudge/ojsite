package problem

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	// "github.com/gogather/com/log"
)

type ProblemDetailController struct {
	controllers.BaseController
}

func (this *ProblemDetailController) Get() {
	id, err := this.GetInt("id")

	if err != nil {
		id = 0
	}

	title := this.Ctx.Input.Param(":title")

	pro := models.Problem{}
	pro, err = pro.GetProblem(id, title)
	if err != nil {
		this.Abort("404")
		return
	}

	this.Data["problem_id"] = pro.Id
	this.Data["problem_title"] = pro.Title
	this.Data["problem_type"] = pro.Type
	this.Data["problem_description"] = pro.Description
	this.Data["problem_pre_code"] = pro.PreCode
	this.Data["problem_tags"] = pro.Tags

	this.TplNames = "problem/detail.tpl"
}
