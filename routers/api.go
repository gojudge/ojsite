package routers

import (
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers/api"
)

func Api() {
	beego.Router("/api/user/username/exist", &api.UsernameExistController{})
	beego.Router("/api/user/email/exist", &api.UserEmailExistController{})
	beego.Router("/api/user/nickname/exist", &api.UserNicknameExistController{})
	beego.Router("/api/problem/list", &api.ProblemListController{})
	beego.Router("/api/problem/title/exist", &api.ProblemTitleExistController{})
}
