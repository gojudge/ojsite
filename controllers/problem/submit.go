package problem

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"strconv"
)

// submit problem
// uri:/problem/submit
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
		this.ServeJson()
		return
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

// get submission status
// uri:/problem/submit/status
type ProblemSubmitStatusController struct {
	controllers.BaseController
}

func (this *ProblemSubmitStatusController) Get() {
	sbid, err := this.GetInt("sbid")

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "get sbid failed",
			"refer":  nil,
		}
		this.ServeJson()
		return
	}

	sub := models.Submissions{}
	status, err := sub.GetSubmissionStatus(sbid)

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"sbid":   sbid,
			"status": nil,
			"msg":    "get status failed",
			"refer":  nil,
		}

	} else {
		this.Data["json"] = map[string]interface{}{
			"result": true,
			"sbid":   sbid,
			"status": status,
			"refer":  nil,
		}
	}

	this.ServeJson()

}
