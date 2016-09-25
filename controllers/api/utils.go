package api

import (
	"github.com/gogather/com/log"
	"github.com/gojudge/ojsite/utils"
	"github.com/labstack/echo"
	"net/http"
)

// parse markdown
func Markdown(c echo.Context) error {
	content := c.FormValue("content")

	log.Blueln(content)
	rst := utils.Markdown2HTML(content)

	res := map[string]interface{}{
		"result":  true,
		"msg":     "success",
		"preview": rst,
		"refer":   nil,
	}

	return c.JSON(http.StatusOK, res)
}
