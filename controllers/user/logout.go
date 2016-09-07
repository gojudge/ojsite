package user

import (
	"github.com/gojudge/ojsite/controllers"
)

type LogoutController struct {
	controllers.BaseController
}

func (this *LogoutController) Get() {
	this.DelSession("username")
	this.DelSession("level")

	this.Redirect("/", 302)
}
