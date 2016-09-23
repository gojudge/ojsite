package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/ojsite/models"
	"github.com/gojudge/ojsite/routers"
	"github.com/gojudge/ojsite/service"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := routers.Init()

	if global.Config == nil {
		log.Info("没有配置文件")
		var err error
		global.Config, err = service.ConfigLoad()
		if err != nil {
			log.Error("读取配置文件错误")
		}
	}
	if service.ConfigIsInit() && !models.EngineInit() {
		service.DBGetEngineWithConfig()
		models.SetEngine(service.Engine)
	}

	e.Run(standard.New(":8088"))

}
