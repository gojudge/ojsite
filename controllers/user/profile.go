package user

import (
	// "github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
)

type ProfileController struct {
	controllers.BaseController
}

func (this *ProfileController) Get() {
	username := this.Ctx.Input.Param(":username")

	if this.Data["userIs"] == "guest" {
		user := &models.User{}
		u, _ := user.GetUser(0, username, "", "")

		this.Data["title"] = u.Nickname + this.Lang("title_profile")
		this.TplNames = "user/profile.tpl"
	} else {
		this.Data["title"] = this.Data["nickname"].(string) + this.Lang("title_user_pannel")
		this.TplNames = "user/pannel.tpl"
	}

}

func (this *ProfileController) Post() {

}