package user

import (
	"fmt"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
)

type RegistorController struct {
	controllers.BaseController
}

func (this *RegistorController) Get() {
	this.Data["title"] = this.Lang("title_registor")
	this.TplNames = "user/registor.tpl"
}

func (this *RegistorController) Post() {
	user := models.User{}

	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	nickname := this.GetString("nickname")

	id, err := user.Registor(username, password, email, nickname)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    fmt.Sprintf("%s", err),
			"refer":  nil,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    fmt.Sprintf("successfully registored, id [%d]", id),
			"refer":  nil,
		}
	}

	this.ServeJson()
}
