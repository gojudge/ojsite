package models

import (
	"github.com/astaxie/beego/orm"
)

type OAuth struct {
	Id    int
	Uid   int
	Type  string
	Token string
}

func (this *OAuth) TableName() string {
	return "oauth"
}

// github login
func (this *OAuth) GithubLogin(token string, oauthType string) (bool, User) {
	o := orm.NewOrm()
	var oa OAuth
	var user User
	var err error

	oa.Token = token
	oa.Type = oauthType
	err = o.QueryTable("oauth").Filter("type", oauthType).Filter("token", token).One(&oa)
	if err != nil {
		return false, user
	}

	err = o.QueryTable("user").Filter("id", oa.Uid).One(&user)
	if err != nil {
		return false, user
	} else if user.Id == 0 {
		return false, user
	} else {
		return true, user
	}
}

// github oauth binding
func (this *OAuth) GithubBind(token string, oauthType string, uid int) bool {
	o := orm.NewOrm()
	var oa OAuth
	var err error

	oa.Uid = uid
	oa.Token = token
	oa.Type = oauthType

	_, err = o.Insert(&oa)
	if err == nil {
		return true
	} else {
		return false
	}
}
