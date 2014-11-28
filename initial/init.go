package initial

import (
	"github.com/duguying/ojsite/routers"
)

// 初始化
func Initialize() {
	InitLang()
	InitTplFunc()
	InitSql()

	routers.Init()
}
