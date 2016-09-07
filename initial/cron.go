package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/gogather/com/log"
	"github.com/gojudge/ojsite/task"
)

func InitCron() {
	tk := toolbox.NewTask("judger", "0/30 * * * * *", task.CheckJudger)

	if beego.AppConfig.String("runmode") == "dev" {
		err := tk.Run()
		if err != nil {
			log.Warnln("[Run Task Failed]")
			log.Warnln(err)
		}
	}

	toolbox.AddTask("judger", tk)
	toolbox.StartTask()
}
