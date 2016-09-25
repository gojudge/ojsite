package problem

import (
	"github.com/gojudge/ojsite/middleware"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	// "github.com/gogather/com/log"
	"net/http"
	"strconv"
)

func ProblemDetail(c echo.Context) error {
	var err error
	idstr := c.FormValue("id")

	id, _ := strconv.Atoi(idstr)

	if err != nil {
		id = 0
	}

	title := c.Param("title")

	pro := models.Problem{}
	pro, err = pro.GetProblem(id, title)
	if err != nil {
		//this.Abort("404")
		return err
	}

	res := c.Get("res").(map[string]interface{})
	cc := c.(*middleware.OJContext)
	res["title"] = cc.Trans("problem") + " - " + pro.Title

	res["problem_id"] = pro.Id
	res["problem_title"] = pro.Title
	res["problem_type"] = pro.Type
	res["problem_description"] = pro.Description
	res["problem_pre_code"] = pro.PreCode
	res["problem_tags"] = pro.Tags

	return c.Render(http.StatusOK, "problem/detail.html", res)
}
