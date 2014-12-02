package models

import (
	"errors"
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
	if len(userName) <= 0 || len(password) < 5 {
		return 0, errors.New("check you form, please.")
	}

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

// get user by `id` or `username` or `email` or `nickname`
func GetUser(id int, username string, email string, nickname string) (User, error) {
	o := orm.NewOrm()
	var user User
	var err error

	if id > 0 {
		user.Id = id
		err = o.Read(&user, "Id")
	} else if len(username) > 0 {
		user.Username = username
		err = o.Read(&user, "Username")
	} else if len(email) > 0 {
		user.Email = email
		err = o.Read(&user, "Email")
	} else if len(nickname) > 0 {
		user.Nickname = nickname
		err = o.Read(&user, "Nickname")
	} else {
		return user, errors.New("at least one parm")
	}

	return user, err
}
