package teach

import (
	// "github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
)

// /teach
type TeachController struct {
	controllers.BaseController
}

func (this *TeachController) Get() {
	this.Data["title"] = this.Lang("title_teach")
	this.TplName = "teach/teach.tpl"
}

func (this *TeachController) Post() {
	this.Data["title"] = this.Lang("title_teach")
	this.TplName = "teach/teach.tpl"
}
