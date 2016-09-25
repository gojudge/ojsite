package initial

import (
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/session"
)

func InitializeSession() {
	path := global.Config.MustValue("session", "path", "tmp")
	ss := session.NewSession(path)
	global.Sessions = &ss
}
