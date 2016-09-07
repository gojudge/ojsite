package controllers

import (
	"fmt"
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/ojsite/service"
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	//return c.String(http.StatusOK, "welcome")
	if global.Config == nil {
		fmt.Println("config is nil")
		var err error
		global.Config, err = service.ConfigLoad()
		if err != nil {
			Install(c)
			fmt.Println("install page")
			return nil
		}
	}
	return c.Render(http.StatusOK, "index1.html", nil)
}

func Install(c echo.Context) error {
	return c.Render(http.StatusOK, "install.html", nil)
}
