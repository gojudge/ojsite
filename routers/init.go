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

	e.Get("/", controllers.Index)

	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
