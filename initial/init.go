package initial

import (
	"github.com/duguying/ojsite/routers"
)

// 初始化
func Initialize() {
	InitErrorHandler()
	InitLang()
	InitTplFunc()
	InitSql()

	routers.Init()
}
