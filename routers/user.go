package routers

import (
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers/user"
)

func User() {
	beego.Router("/login", &user.LoginController{})
	beego.Router("/logout", &user.LogoutController{})
	beego.Router("/register", &user.RegisterController{})
	beego.Router("/oauth/github", &user.OAuthController{})
	beego.Router("/oauth/osc", &user.OAuthOSCController{})
	beego.Router("/student/verify", &user.StudentVerityController{})
	beego.Router("/user/setting/profile", &user.SettingProfileController{})
	beego.Router("/user/setting/pwd", &user.SettingPwdController{})
	beego.Router("/user/setting/bind", &user.SettingBindController{})
	beego.Router("/:username", &user.ProfileController{})
}
