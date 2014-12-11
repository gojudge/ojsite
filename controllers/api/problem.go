package api

import (
	// "github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
)

// get problem list
type ProblemListController struct {
	controllers.BaseController
}

func (this *ProblemListController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJson()
}

// check problem title exist
type ProblemTitleExistController struct {
	controllers.BaseController
}

func (this *ProblemTitleExistController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJson()
}
