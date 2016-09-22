package routers

import (
	"github.com/echo-contrib/pongor"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Debug()

	r := pongor.GetRenderer()
	e.SetRenderer(r)

	//e.Use(middleware.Prepare)

	e.Static("static", "static/dist")
	e.Static("staticasset", "static/asset")

	e.Get("/", controllers.Index)

	e.Use(middleware.UserCheck)

	rInstall := e.Group("/install")
	{
		rInstall.Post("/do_submit", controllers.InstallDoSubmit)
		rInstall.Get("/", controllers.InstallIndex)
	}

	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
