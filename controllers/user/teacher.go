package user

import (
	"github.com/gojudge/ojsite/controllers"
)

type TeacherPannelController struct {
	controllers.BaseController
}

func (this *TeacherPannelController) Get() {
	this.Forbbiden("not", "teacher")
	this.Data["title"] = this.Lang("title_teacher_pannel")
	this.TplName = "teacher/teacher_pannel.tpl"
}
