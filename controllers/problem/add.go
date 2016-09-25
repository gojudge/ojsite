package problem

import (
	"github.com/gojudge/ojsite/middleware"
	"github.com/labstack/echo"
	"net/http"
)

func ProblemAdd(c echo.Context) error {
	res := c.Get("res").(map[string]interface{})
	cc := c.(*middleware.OJContext)
	res["title"] = cc.Trans("title_add_problem")
	return c.Render(http.StatusOK, "problem/add.html", res)
}

func DoProblemAdd(c echo.Context) error {
	res := c.Get("res").(map[string]interface{})
	cc := c.(*middleware.OJContext)
	res["title"] = cc.Trans("title_add_problem")
	return c.JSON(http.StatusOK, res)
}
