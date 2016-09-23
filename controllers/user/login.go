package user

import (
	"fmt"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"net/http"
)

func Login(c echo.Context) error {
	//this.Forbbiden("not", "guest")

	res := make(map[string]interface{})
	res["title"] = "登录"
	//res["github_client_id"] = global.Config.GetValue("env", "github_client_id")
	return c.Render(http.StatusOK, "user/login.html", res)
}

func DoLogin(c echo.Context) error {
	//this.Forbbiden("not", "guest")
	user := models.User{}
	username := c.FormValue("username")
	password := c.FormValue("password")
	res := make(map[string]interface{})

	if result, lev := user.Login(username, password); result {
		user, err := user.GetUser(0, username, "", "")
		if err != nil {
			res = map[string]interface{}{
				"result": false,
				"msg":    "login failed, get user info failed",
				"refer":  nil,
			}
		} else {
			cookie := new(echo.Cookie)
			cookie.SetName("username")
			cookie.SetValue(username)
			c.SetCookie(cookie)
			cookie1 := new(echo.Cookie)
			cookie1.SetName("userid")
			cookie1.SetValue(fmt.Sprintf("%d", user.Id))
			c.SetCookie(cookie1)
			cookie2 := new(echo.Cookie)
			cookie2.SetName("level")
			cookie2.SetValue(lev)
			c.SetCookie(cookie2)

			res = map[string]interface{}{
				"result": true,
				"msg":    "login success",
				"refer":  nil,
			}
		}

	} else {
		res = map[string]interface{}{
			"result": false,
			"msg":    "login failed",
			"refer":  nil,
		}
	}

	return c.JSON(http.StatusOK, res)
}
