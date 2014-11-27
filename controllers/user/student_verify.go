package user

import (
	"github.com/duguying/ojsite/controllers"
)

type StudentVerityController struct {
	controllers.BaseController
}

func (this *StudentVerityController) Get() {
	this.Data["title"] = this.Lang("title_student_verify")
	this.TplNames = "user/student_verify.tpl"
}
