package user

import (
	"github.com/duguying/ojsite/controllers"
)

type RegistorController struct {
	controllers.BaseController
}

func (this *RegistorController) Get() {
	this.Data["title"] = this.Lang("title_registor")
	this.TplNames = "user/registor.tpl"
}
