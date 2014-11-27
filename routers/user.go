package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers/user"
)

func User() {
	beego.Router("/login", &user.LoginController{})
	beego.Router("/registor", &user.RegistorController{})
	beego.Router("/oauth/github", &user.OAuthController{})
	beego.Router("/student/verify", &user.StudentVerityController{})
}
