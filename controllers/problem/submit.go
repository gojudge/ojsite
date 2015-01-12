package problem

import (
	// "github.com/duguying/judger/client"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
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

	user := this.GetSession("user").(models.User)

	sub := models.Submissions{}
	sub.Add(pid, user.Id, ptype, language, code, "default")

}
