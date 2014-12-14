package problem

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"github.com/gogather/com/log"
)

type ProblemListController struct {
	controllers.BaseController
}

func (this *ProblemListController) Get() {
	pro := &models.Problem{}
	problems, hasNext, _, _ := pro.ListProblem(1, 10, "public", "ok")

	top10, _ := pro.GetTop10()
	log.Blueln("[top 10]", top10)

	tag := &models.Tags{}
	tagList, _ := tag.TagList()
	log.Blueln("[tags]", tagList)

	this.Data["title"] = this.Lang("title_problem_list")

	this.Data["problems"] = problems
	this.Data["title"] = this.Lang("title_problems")
	this.Data["has_next"] = hasNext

	this.Data["top10"] = top10
	this.Data["tag_list"] = tagList

	this.TplNames = "problem/list.tpl"
}
