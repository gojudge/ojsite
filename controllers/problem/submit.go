package problem

import (
	"github.com/duguying/judger/client"
	"github.com/duguying/ojsite/controllers"
	// "github.com/duguying/ojsite/models"
	// "github.com/gogather/com/log"
)

type ProblemSubmitController struct {
	controllers.BaseController
}

func (this *ProblemSubmitController) Post() {
	// pid, err := this.GetInt("pid")
	// language := this.GetString("language")
	// ptype := this.GetString("type")
	// code := this.GetString("code")

	// sub := models.Submissions{}
	// sub.Add(pid, uid, code, judger)

	test := `{"action":"login","password":"123456789"}#`
	client.J.Request(test)
}
