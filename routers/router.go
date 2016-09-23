package routers

import (
	"github.com/echo-contrib/pongor"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/controllers/user"
	"github.com/gojudge/ojsite/middleware"
	"github.com/labstack/echo"
	mm "github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Debug()

	r := pongor.GetRenderer(pongor.PongorOption{
		Reload: true,
	})
	e.SetRenderer(r)

	//e.Use(middleware.Prepare)

	e.Static("static", "static/dist")
	e.Static("staticasset", "static/asset")

	e.Get("/", controllers.Index)

	e.Use(middleware.UserCheck)

	e.Use(mm.CSRF())

	rInstall := e.Group("/install")
	{
		rInstall.Post("/do_submit", controllers.InstallDoSubmit)
		rInstall.Get("/", controllers.InstallIndex)
	}

	e.Get("/login", user.Login)
	e.Post("/do_login", user.DoLogin)

	e.Get("/register", user.Register)

	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
