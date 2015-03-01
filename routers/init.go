package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})

	User()
	Problem()
	Teacher()
	Api()
	Teach()
}
