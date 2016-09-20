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
			fmt.Println("load config file error")
			fmt.Println(global.Config)
		}
	}

	dburl, _ := global.Config.GetValue("database", "dburl")
	fmt.Println("dburl is %s\n", dburl)
	if dburl == "" {
		fmt.Println("go to install page")
		return InstallIndex(c)
	}

	return c.Render(http.StatusOK, "index.html", nil)
}

func InstallIndex(c echo.Context) error {
	fmt.Println("here it is install page")
	return c.Render(http.StatusOK, "install.html", nil)
}

// InstallDoSubmit to init a new database
func InstallDoSubmit(c echo.Context) error {
	dburl := c.FormValue("dburl")
	dbport := c.FormValue("dbport")
	dbuser := c.FormValue("dbuser")
	dbpwd := c.FormValue("dbpwd")
	dbname := c.FormValue("dbname")
	adminuser := c.FormValue("admin-user")
	adminpwd := c.FormValue("adminpwd")
	err := service.DBInit(dburl, dbport, dbuser, dbpwd, dbname, adminuser, adminpwd)
	res := make(map[string]interface{})
	res["success"] = true
	if err != nil {
		res["success"] = false
		res["msg"] = err
		return c.JSON(http.StatusOK, res)
	}
	return c.JSON(http.StatusOK, res)
}
