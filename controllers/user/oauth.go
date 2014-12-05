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
	this.Forbbiden("login")

	code := this.GetString("code")
	clientId := beego.AppConfig.String("github_client_id")
	clientSecret := beego.AppConfig.String("github_client_secret")

	oauthGithub := &oauth.GithubOAuth{}
	token, _, err := oauthGithub.GetData(clientId, clientSecret, code)

	if err != nil {
		this.Ctx.WriteString("Response Error! ")
		return
	}

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

}
