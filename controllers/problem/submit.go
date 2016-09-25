package problem

import (
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// ProblemSubmit to submit problem to database
// uri:/problem/submit
func ProblemSubmit(c echo.Context) error {
	pidStr := c.FormValue("pid")
	pid, err := strconv.Atoi(pidStr)

	if err != nil {
		res := map[string]interface{}{
			"result": false,
			"msg":    "get pid failed",
			"refer":  nil,
		}
		return c.JSON(http.StatusOK, res)
	}

	language := c.FormValue("language")
	ptype := c.FormValue("type")
	code := c.FormValue("code")

	userid := c.Get("userID")
	if userid == nil {
		res := map[string]interface{}{
			"result": false,
			"msg":    "unknown userid",
			"debug":  err,
			"refer":  nil,
		}
		return c.JSON(http.StatusOK, res)
	}

	useridInt := userid.(int)
	sub := models.Submissions{}
	id, err := sub.Add(pid, useridInt, ptype, language, code, "default")

	var res map[string]interface{}
	if err != nil {
		res = map[string]interface{}{
			"result": false,
			"msg":    "add task failed",
			"debug":  err,
			"refer":  nil,
		}
	} else {
		res = map[string]interface{}{
			"result": true,
			"msg":    "add task success",
			"id":     id,
			"refer":  nil,
		}
	}
	return c.JSON(http.StatusOK, res)
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
		this.ServeJSON()
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

	this.ServeJSON()

}
