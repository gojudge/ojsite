package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	clog "github.com/gogather/com/log"
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/ojsite/initial"
	"github.com/gojudge/ojsite/models"
	"github.com/gojudge/ojsite/routers"
	"github.com/gojudge/ojsite/service"
	"github.com/labstack/echo/engine/standard"
	"os"
	"strconv"
	"strings"
)

const (
	PORT = 8088
)

func main() {
	e := routers.Init()
	preLoadConfig()
	parseArgs()
	initial.InitLang()
	initial.InitializeSession()
	clog.Pinkln(global.Config)
	port, err := global.Config.Int64("", "port")
	if err != nil {
		port = PORT
	}
	e.Run(standard.New(p(port)))
}

func parseArgs() {
	args := os.Args
	if len(args) > 1 {
		expectPort := args[1]
		portArgs := strings.Split(expectPort, "=")
		clog.Greenln(portArgs)
		if len(portArgs) > 1 && portArgs[0] == "--port" {
			port, err := strconv.ParseInt(portArgs[1], 10, 64)
			if err == nil {
				global.Config.SetValue("config", "port", s(port))
			} else {
				global.Config.SetValue("config", "port", s(PORT))
			}
		}
	} else {
		global.Config.SetValue("config", "port", s(PORT))
	}
}

func p(port int64) string {
	return fmt.Sprintf(":%d", port)
}

func s(port int64) string {
	return fmt.Sprintf("%d", port)
}

func preLoadConfig() {
	//return c.String(http.StatusOK, "welcome")
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
}
