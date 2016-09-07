package main

import (
	"github.com/astaxie/beego"
	"github.com/gojudge/ojsite/initial"
)

const (
	APP_VER = "0.0.2.0205"
)

func init() {
	initial.AppVer = APP_VER
}

func main() {
	initial.Initialize()
	beego.Run()
}
