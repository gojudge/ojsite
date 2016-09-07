package user

import (
	// "github.com/astaxie/beego"
	"github.com/gogather/com/log"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/models"
)

type ProfileController struct {
	controllers.BaseController
}

func (this *ProfileController) Get() {
	username := this.Ctx.Input.Param(":username")

	ip := this.Ctx.Input.IP()
	log.Blueln(ip)

	if this.Data["userIs"] == "guest" {
		user := &models.User{}
		u, _ := user.GetUser(0, username, "", "")

		this.Data["title"] = u.Nickname + this.Lang("title_profile")

	} else {
		this.Data["title"] = this.Data["nickname"].(string) + this.Lang("title_profile")
	}

	this.TplName = "user/profile.tpl"

}

func (this *ProfileController) Post() {
	this.Data["json"] = map[string]interface{}{
		"result": false,
		"msg":    "only get method is valid",
	}

	this.ServeJSON()
}
