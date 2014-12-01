package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

func Trace(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fmt.Printf(format, v...)
	log.Trace(format, v...)
}

func Warn(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fmt.Printf(format, v...)
	log.Warn(format, v...)
}
