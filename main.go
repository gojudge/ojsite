package main

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/initial"
)

func main() {
	initial.Initialize()
	beego.Run()
}
