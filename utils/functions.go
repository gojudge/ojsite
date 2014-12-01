package utils

import (
	"github.com/astaxie/beego/logs"
)

func Trace(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	log.Trace(format, v...)
}

func Warn(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	log.Warn(format, v...)
}
