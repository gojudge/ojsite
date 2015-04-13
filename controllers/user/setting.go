package user

import (
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
	// "github.com/gogather/com/log"
)

type SettingController struct {
	controllers.BaseController
}

func (this *SettingController) Get() {

	this.Data["title"] = this.Lang("title_user_pannel")

	this.TplNames = "user/pannel.tpl"
}
