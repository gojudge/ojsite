package init

import (
	"github.com/duguying/ojsite/routers"
)

// 初始化
func init() {
	InitLang()
	InitTplFunc()

	routers.Init()
}
