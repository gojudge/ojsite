package controllers

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/ojsite/service"
	"github.com/labstack/echo"
	"net/http"
)

// Index index page
func Index(c echo.Context) error {
	//return c.String(http.StatusOK, "welcome")
	if global.Config == nil {
		log.Info("没有配置文件")
		var err error
		global.Config, err = service.ConfigLoad()
		if err != nil {
			log.Error("读取配置文件错误")
		}
	}

	dburl, _ := global.Config.GetValue("database", "dburl")
	fmt.Println("dburl is %s\n", dburl)
	if dburl == "" {
		log.Info("准备安装")
		return InstallIndex(c)
	}

	return c.Render(http.StatusOK, "index.html", nil)
}

// InsallIndex install page
func InstallIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "install.html", nil)
}

// InstallDoSubmit to init a new database
func InstallDoSubmit(c echo.Context) error {
	dburl := c.FormValue("dburl")
	dbport := c.FormValue("dbport")
	dbuser := c.FormValue("dbuser")
	dbpwd := c.FormValue("dbpwd")
	dbname := c.FormValue("dbname")
	adminuser := c.FormValue("adminuser")
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
