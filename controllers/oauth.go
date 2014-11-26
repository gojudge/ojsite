package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gogather/oauth"
)

type OAuthController struct {
	BaseController
}

func (this *OAuthController) Get() {

	code := this.GetString("code")
	clientId := beego.AppConfig.String("github_client_id")
	clientSecret := beego.AppConfig.String("github_client_secret")

	oauthGithub := &oauth.GithubOAuth{}
	json, err := oauthGithub.GetData(clientId, clientSecret, code)

	if err != nil {
		this.Ctx.WriteString("Response Error! ")
		return
	}
	data := json.(map[string]interface{})

	this.Data["login"] = data["login"].(string)
	this.Data["avatar_url"] = data["avatar_url"].(string)
	this.Data["name"] = data["name"].(string)

	this.TplNames = "user/oauth.tpl"
}
