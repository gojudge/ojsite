package user

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/duguying/ojsite/controllers"
	"github.com/duguying/ojsite/models"
	"github.com/gogather/com/log"
	"github.com/gogather/oauth"
)

type OAuthController struct {
	controllers.BaseController
}

func (this *OAuthController) Get() {
	code := this.GetString("code")
	clientId := beego.AppConfig.String("github_client_id")
	clientSecret := beego.AppConfig.String("github_client_secret")

	oauthGithub := &oauth.GithubOAuth{}
	token, json, err := oauthGithub.GetData(clientId, clientSecret, code)

	if err != nil {
		this.Ctx.WriteString("Response Error! ")
		return
	} else {
		log.Blueln(json)
	}

	if this.Data["userIs"] == "guest" {

		user := models.User{}
		result, usr := user.GithubLogin(token)
		if result {
			log.Blueln("github login success.")
			this.SetSession("username", usr.Username)
			this.SetSession("level", usr.Level)
			this.Redirect("/", 302)
		} else {
			log.Warnln("you have not register or bindding you github account.")
			this.Redirect("/register", 302)
		}

	} else {
		// have login, binding github oauth
		data := json.(map[string]interface{})
		avatar := data["avatar_url"].(string)
		username := data["login"].(string)
		this.Data["title"] = "github绑定确认"
		this.Data["github_avatar"] = avatar
		this.Data["goj_avatar"] = avatar
		this.Data["token"] = token
		this.Data["github_username"] = username

		this.TplNames = "user/github_binding_confirm.tpl"
	}

}

func (this *OAuthController) Post() {
	this.Forbbiden("guest")

	password := this.GetString("password")
	token := this.GetString("token")
	username := this.Data["username"].(string)
	user := models.User{}
	result, _ := user.Login(username, password)
	if result {
		if user.GithubBind(token, username) {
			this.Data["json"] = map[string]interface{}{
				"result": true,
				"msg":    "binding success",
				"refer":  nil,
			}
		} else {
			this.Data["json"] = map[string]interface{}{
				"result": false,
				"msg":    "binding failed",
				"refer":  nil,
			}
		}

	} else {
		this.Data["json"] = map[string]interface{}{
			"result": false,
			"msg":    "wrong password.",
			"refer":  nil,
		}
	}

	this.ServeJson()

}
