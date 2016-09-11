package global

import (
	"github.com/Unknwon/goconfig"
	"github.com/go-xorm/xorm"
)

var (
	Config *goconfig.ConfigFile
	Engine *xorm.Engine
)
