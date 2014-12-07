package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/duguying/ojsite/utils"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
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
// return: int64 id
//         error if failed
func (this *User) Register(userName string, password string, email string, nickName string) (int64, error) {
	if len(userName) <= 0 || len(password) < 5 || len(email) <= 0 {
		return 0, errors.New("check you form, please.")
	}

	if len(nickName) <= 0 {
		nickName = userName + com.RandString(5) // gen the default nickname
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

// user login
// return: bool if login
//         string user level if exist
func (this *User) Login(userName string, password string) (bool, string) {
	if len(userName) <= 0 || len(password) <= 0 {
		return false, ""
	}

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
// return: User if successfully get
//         error if getting failed, and User is empty
func (this *User) GetUser(id int, username string, email string, nickname string) (User, error) {
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

// get avatar
func (this *User) GetAvatar(id int, username string, email string, nickname string) (string, error) {
	user, err := this.GetUser(id, username, email, nickname)
	if nil == err {
		return beego.AppConfig.String("avatar") + com.Md5(user.Email), err
	} else {
		log.Warnln("GetAvatar Failed.", err)
		return "", err
	}
}
