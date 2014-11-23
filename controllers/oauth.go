package controllers

import (
	"github.com/gogather/oauth"
)

type OAuthController struct {
	BaseController
}

func (this *OAuthController) Post() {
	oauthGithub := &oauth.GithubOAuth{}
	json, err := oauthGithub.GetData()
	if err != nil {
		this.Ctx.WriteString("Response Error!")
	}
	data := json.(map[string]interface{})

	this.Data["login"] = data["login"].(string)
	this.Data["avatar_url"] = data["avatar_url"].(string)
	this.Data["name"] = data["name"].(string)
}
