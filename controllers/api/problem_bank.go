package api

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"strconv"
)

// get problem list
type ProblemBankAcceptController struct {
	controllers.BaseController
}

func (this *ProblemBankAcceptController) Get() {
	// get id
	id, err := this.GetInt("id")
	if nil != err || id < 0 {
		id = 0
	}

	s := this.Ctx.Input.Param(":id")
	pageParm, err := strconv.Atoi(s)
	if nil != err || pageParm < 0 {
		pageParm = 0
	} else {
		id = pageParm
	}

	if 0 == id {
		id = 1
	}

	// accept via id
	prob := models.ProblemBank{}
	err = prob.AcceptProblem(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "accept problem failed",
			"refer":  nil,
			"debug":  err,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "problem accepted",
			"refer":  nil,
		}
	}

	this.ServeJson()
}
