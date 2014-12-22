package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

func Init() {
	beego.Router("/", &controllers.MainController{})

	User()
	Problem()
	Teacher()
	Api()
	Teach()
}
