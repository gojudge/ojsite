package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"time"
)

type User struct {
	Id           int
	Username     string
	Password     string
	Salt         string
	Email        string
	Level        string
	RegistorTime time.Time
	Nickname     string
}

// user registor
func (this *User) Registor(userName string, password string, email string, nickName string) (int64, error) {
	o := orm.NewOrm()
	var user User
	user.Username = userName
	user.Salt = com.RandString(7)
	user.Password = com.Md5(password + user.Salt)
	user.Email = email
	user.Level = "user"
	user.Nickname = nickName

	return o.Insert(&user)
}
