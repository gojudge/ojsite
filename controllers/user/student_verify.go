package user

import (
	"github.com/gojudge/ojsite/controllers"
)

type StudentVerityController struct {
	controllers.BaseController
}

func (this *StudentVerityController) Get() {
	this.Data["title"] = this.Lang("title_student_verify")
	this.TplName = "user/student_verify.tpl"
}
