package init

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/duguying/ojsite/controllers"
)

func i18nGetString(key string) string {
	return i18n.Tr(controllers.LANG, key)
}

func InitTplFunc() {
	beego.AddFuncMap("i18n", i18nGetString)
}
