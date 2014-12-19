package main

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/initial"
)

const (
	APP_VER = "0.0.1.1220"
)

func init() {
	initial.AppVer = APP_VER
}

func main() {
	initial.Initialize()
	beego.Run()
}
