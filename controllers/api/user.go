package api

import (
	// "github.com/astaxie/beego"
	"github.com/gojudge/ojsite/controllers"
	// "github.com/gojudge/ojsite/models"
)

// check username exist
type UsernameExistController struct {
	controllers.BaseController
}

func (this *UsernameExistController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJSON()
}

// check user email exist
type UserEmailExistController struct {
	controllers.BaseController
}

func (this *UserEmailExistController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJSON()
}

// check user nickname exist
type UserNicknameExistController struct {
	controllers.BaseController
}

func (this *UserNicknameExistController) Get() {
	this.Data["json"] = map[string]interface{}{
		"result": true,
		"msg":    "get list success",
		"list":   nil,
		"refer":  nil,
	}

	this.ServeJSON()
}
