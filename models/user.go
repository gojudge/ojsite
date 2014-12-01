package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/duguying/ojsite/utils"
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
func (this *User) Register(userName string, password string, email string, nickName string) (int64, error) {
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

func Login(userName string, password string) (bool, string) {
	o := orm.NewOrm()
	var user User
	user.Username = userName

	err := o.Read(&user, "Username")

	if err != nil {
		utils.Trace("查询不到")
	} else {
		pwd := com.Md5(password + user.Salt)
		utils.Trace("[pwd] %s\n[password] %s", pwd, user.Password)
		if user.Password == pwd {
			return true, user.Level
		} else {
			return false, ""
		}
	}
	return false, ""
}
