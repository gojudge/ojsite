package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers/problem"
)

func Problem() {
	beego.Router("/problem", &problem.ProblemsController{})
}
