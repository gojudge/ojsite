package global

import (
	"github.com/Unknwon/goconfig"
	"github.com/go-xorm/xorm"
	"github.com/gojudge/session"
)

var (
	Config   *goconfig.ConfigFile
	Engine   *xorm.Engine
	Sessions *session.Session
)
