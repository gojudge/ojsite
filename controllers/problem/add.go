package problem

import (
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
	// "github.com/gogather/com/log"
)

type AddProblemController struct {
	controllers.BaseController
}

func (this *AddProblemController) Get() {
	this.Data["title"] = this.Lang("title_add_problem")

	this.TplNames = "problem/add.tpl"
}
