package main

import (
	"github.com/gojudge/ojsite/routers"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := routers.Init()

	e.Run(standard.New(":8088"))

}
