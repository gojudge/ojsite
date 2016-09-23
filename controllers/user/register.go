package user

import (
	"fmt"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"net/http"
)

func Register(c echo.Context) error {
	//this.Forbbiden("not", "guest")

	res := make(map[string]interface{})
	res["title"] = "title_register"
	return c.Render(http.StatusOK, "user/register.html", res)
}

func Post(c echo.Context) error {
	//this.Forbbiden("not", "guest")

	user := models.User{}

	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	nickname := c.FormValue("nickname")

	id, err := user.Register(username, password, email, nickname)
	res := make(map[string]interface{})
	if err != nil {
		res = map[string]interface{}{
			"result": false,
			"msg":    fmt.Sprintf("%s", err),
			"refer":  nil,
		}
	} else {
		res = map[string]interface{}{
			"result": true,
			"msg":    fmt.Sprintf("successfully registered, id [%d]", id),
			"refer":  nil,
		}
	}
	return c.JSON(http.StatusOK, res)
}
