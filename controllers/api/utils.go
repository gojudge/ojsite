package api

import (
	"github.com/gogather/com/log"
	"github.com/gojudge/ojsite/controllers"
	"github.com/gojudge/ojsite/utils"
)

// parse markdown
type MarkdownController struct {
	controllers.BaseController
}

func (this *MarkdownController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": false,
		"msg":    "only post method support",
		"refer":  nil,
	}

	this.ServeJSON()
}

func (this *MarkdownController) Post() {
	content := this.GetString("content")

	log.Blueln(content)
	rst := utils.Markdown2HTML(content)

	this.Data["json"] = map[string]interface{}{
		"result":  true,
		"msg":     "success",
		"preview": rst,
		"refer":   nil,
	}

	this.ServeJSON()
}
