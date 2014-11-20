package main

import (
	"github.com/astaxie/beego"
	_ "github.com/duguying/ojsite/init"
)

const (
	APP_VER = "0.0.1.1121"
)

func main() {
	beego.Run()
}
