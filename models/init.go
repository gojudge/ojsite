package models

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	//log.SetFormatter(&log.TextFormatter{})
}

func SetEngine(e *xorm.Engine) {
	engine = e
}

func EngineInit() bool {
	return engine != nil
}
