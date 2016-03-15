package user

import (
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
	// "github.com/gogather/com/log"
)

// setting profile
type SettingProfileController struct {
	controllers.BaseController
}

func (this *SettingProfileController) Get() {
	this.Data["title"] = this.Lang("title_user_setting_profile")
	this.Data["nav"] = "profile"

	this.TplName = "user/setting/profile.tpl"
}

// setting password
type SettingPwdController struct {
	controllers.BaseController
}

func (this *SettingPwdController) Get() {
	this.Data["title"] = this.Lang("title_user_setting_pwd")
	this.Data["nav"] = "pwd"

	this.TplName = "user/setting/pwd.tpl"
}

// setting binding
type SettingBindController struct {
	controllers.BaseController
}

func (this *SettingBindController) Get() {
	this.Data["title"] = this.Lang("title_user_setting_bind")
	this.Data["nav"] = "bind"

	this.TplName = "user/setting/bind.tpl"
}
