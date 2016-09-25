package problem

import (
	"github.com/gogather/com/log"
	"github.com/gojudge/ojsite/middleware"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Num  int
	Name string
}

func ListProblems(c echo.Context) error {
	pro := &models.Problem{}
	problems, hasNext, _, _ := pro.ListProblem(1, 10, "public")
	log.Blueln("[problems]", problems)

	top10, _ := pro.GetTop10()
	log.Blueln("[top 10]", top10)

	tag := &models.Tags{}
	tagList, _ := tag.TagList()
	log.Blueln("[tags]", tagList)

	res := c.Get("res").(map[string]interface{})
	cc := c.(*middleware.OJContext)

	res["title"] = cc.Trans("title_problem_list")

	res["problems"] = problems
	res["title"] = cc.Trans("title_problems")
	res["has_next"] = hasNext

	res["top10"] = top10
	res["tag_list"] = tagList

	return c.Render(http.StatusOK, "problem/list.html", res)
}
