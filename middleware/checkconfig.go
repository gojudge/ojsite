package middleware

import (
//"fmt"
//"github.com/gojudge/ojsite/controllers"
//"github.com/gojudge/ojsite/global"
//"github.com/gojudge/ojsite/service"
//"github.com/labstack/echo"
//"strings"
)

/// CheckConfig to check if config file exist and orm engine is exist
/*
func CheckConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() != "/" && c.Path() != "/index" && strings.Contains(c.Path(), "/install/") {
			if global.Config == nil {
				var err error
				global.Config, err = service.ConfigLoad()
				if err != nil {
					//controllers.InstallIndex(c)
					fmt.Println("Please viste index page to install")
					next(c)
				}
			}
		}
		return next(c)
	}
}*/

//func Prepare(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		res := make(map[string]interface{})
//
//	}
//}
