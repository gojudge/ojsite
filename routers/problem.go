package routers

import (
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers/problem"
)

func Problem() {
	beego.Router("/problems", &problem.ProblemListController{})
	beego.Router("/problem/:title", &problem.ProblemDetailController{})
	beego.Router("/problem/add", &problem.AddProblemController{})
	beego.Router("/problem/submit", &problem.ProblemSubmitController{})
	beego.Router("/problem/submit/status", &problem.ProblemSubmitStatusController{})
}
