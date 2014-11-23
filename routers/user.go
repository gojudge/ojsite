package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
)

func User() {
	beego.Router("/login", &controllers.LoginController{})
}
