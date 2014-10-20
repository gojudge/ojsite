package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
