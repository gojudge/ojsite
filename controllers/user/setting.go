package user

import (
	"github.com/gojudge/ojsite/controllers"
	// "github.com/gojudge/ojsite/models"
	"github.com/gojudge/ojsite/middleware"
	"github.com/labstack/echo"
	"net/http"
	// "github.com/gogather/com/log"
)

// setting profile

func SettingProfile(c echo.Context) error {
	cc := c.(*middleware.OJContext)
	res := c.Get("res").(map[string]interface{})
	res["title"] = cc.Trans("title_user_setting_profile")
	res["nav"] = "profile"

	return c.Render(http.StatusOK, "user/setting/profile.html", res)
}

// setting password
func SettingPwd(c echo.Context) error {
	cc := c.(*middleware.OJContext)
	res := c.Get("res").(map[string]interface{})
	res["title"] = cc.Trans("title_user_setting_pwd")
	res["nav"] = "pwd"

	return c.Render(http.StatusOK, "user/setting/pwd.html", res)
}

// setting binding
type SettingBindController struct {
	controllers.BaseController
}

func (this *SettingBindController) Get() {
	this.Data["title"] = this.Lang("title_user_setting_bind")
	this.Data["nav"] = "bind"

	this.TplName = "user/setting/bind.tpl"
}
