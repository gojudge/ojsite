package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers/problem"
)

func Problem() {
	beego.Router("/problems", &problem.ProblemListController{})
	beego.Router("/problem/:title", &problem.ProblemDetailController{})
	beego.Router("/problem/add", &problem.AddProblemController{})
}
