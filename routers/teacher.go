package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers/user"
)

func Teacher() {
	beego.Router("/teacher", &user.TeacherPannelController{})
}
