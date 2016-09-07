package routers

import (
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers/user"
)

func Teacher() {
	beego.Router("/teacher", &user.TeacherPannelController{})
}
