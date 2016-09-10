package routers

import (
	"github.com/echo-contrib/pongor"
	"github.com/gojudge/ojsite/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Debug()

	r := pongor.GetRenderer()
	e.SetRenderer(r)

	e.Static("static", "static/dist")
	e.Static("staticasset", "static/asset")

	e.Get("/", controllers.Index)

	rInstall := e.Group("/install")
	{
		rInstall.Get("/index", controllers.InstallIndex)
		rInstall.Post("/do_submit", controllers.InstallDoSubmit)
	}

	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
