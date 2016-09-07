package user

import (
	"fmt"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/models"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) Get() {
	this.Forbbiden("not", "guest")

	this.Data["title"] = this.Lang("title_register")
	this.TplName = "user/register.tpl"
}

func (this *RegisterController) Post() {
	this.Forbbiden("not", "guest")

	user := models.User{}

	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	nickname := this.GetString("nickname")

	id, err := user.Register(username, password, email, nickname)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    fmt.Sprintf("%s", err),
			"refer":  nil,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    fmt.Sprintf("successfully registered, id [%d]", id),
			"refer":  nil,
		}
	}

	this.ServeJSON()
}
