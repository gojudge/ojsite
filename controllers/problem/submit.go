package problem

import (
	// "github.com/duguying/judger/client"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"strconv"
	// "github.com/gogather/com/log"
)

type ProblemSubmitController struct {
	controllers.BaseController
}

func (this *ProblemSubmitController) Post() {
	pid, err := this.GetInt("pid")

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "get pid failed",
			"refer":  nil,
		}
	}

	language := this.GetString("language")
	ptype := this.GetString("type")
	code := this.GetString("code")

	userid := this.GetSession("userid").(string)
	useridInt, err := strconv.Atoi(userid)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "unknown userid",
			"debug":  err,
			"refer":  nil,
		}
		this.ServeJson()
		return
	}

	sub := models.Submissions{}
	id, err := sub.Add(pid, useridInt, ptype, language, code, "default")

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "add task failed",
			"debug":  err,
			"refer":  nil,
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"msg":    "add task success",
			"id":     id,
			"refer":  nil,
		}
	}
	this.ServeJson()

}
