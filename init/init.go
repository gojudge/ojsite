package init

import (
	"github.com/duguying/ojsite/routers"
)

// 初始化
func init() {
	InitConfig()
	InitLang()
	InitTplFunc()

	routers.Init()
}
