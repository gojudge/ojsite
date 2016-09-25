package routers

import (
	_ "github.com/digitalcrab/pongo2trans"
	"github.com/echo-contrib/pongor"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/controllers/problem"
	"github.com/gojudge/ojsite/controllers/user"
	"github.com/gojudge/ojsite/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.SetDebug(true)

	r := pongor.GetRenderer(pongor.PongorOption{
		Reload: true,
	})
	e.SetRenderer(r)

	//e.Use(middleware.Prepare)

	e.Static("static", "static/dist")
	e.Static("staticasset", "static/asset")

	e.Use(middleware.OJHandle)
	e.Use(middleware.UserCheck)

	e.Get("/", controllers.Index)

	e.Get("/login", user.Login)
	e.Post("/do_login", user.DoLogin)

	e.Get("/register", user.Register)
	e.Post("/do_register", user.DoRegister)

	rInstall := e.Group("/install")
	{
		rInstall.Post("/do_submit", controllers.InstallDoSubmit)
		rInstall.Get("/", controllers.InstallIndex)
	}

	rUser := e.Group("/user")
	{
		rUser.Get("/setting/profile", user.SettingProfile)
		rUser.Get("/setting/pwd", user.SettingPwd)
	}

	rProblems := e.Group("/problems")
	{
		rProblems.Get("/", problem.ListProblems)
		rProblems.Get("/p/:title", problem.ProblemDetail)
		rProblems.Post("/submit", problem.ProblemSubmit)
	}
	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
