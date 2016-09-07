package routers

import (
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers/teach"
)

func Teach() {
	beego.Router("/teach", &teach.TeachController{})
}
