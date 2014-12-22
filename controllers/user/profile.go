package user

import (
	// "github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
)

type ProfileController struct {
	controllers.BaseController
}

func (this *ProfileController) Get() {
	this.Data["title"] = this.Data["nickname"].(string) + this.Lang("title_profile")

	this.TplNames = "user/profile.tpl"
}

func (this *ProfileController) Post() {

}
