package models

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/gogather/com"
	"github.com/gojudge/ojsite/global"
	"github.com/gojudge/ojsite/utils"
	"time"
)

type User struct {
	Id           int
	Username     string
	Password     string
	Salt         string
	Email        string
	Level        string
	RegistorTime time.Time `xorm:"created"`
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

	var user User
	user.Username = userName
	user.Salt = com.RandString(7)
	user.Password = com.Md5(password + user.Salt)
	user.Email = email
	user.Level = "user"
	user.Nickname = nickName

	_, err := engine.Insert(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"msg": err,
		}).Error("添加用户错误")
		return 0, err
	}
	return int64(user.Id), nil
}

// user login
// return: bool if login
//         string user level if exist
func (this *User) Login(userName string, password string) (bool, string) {
	if len(userName) <= 0 || len(password) <= 0 {
		return false, ""
	}

	var user User
	user.Username = userName

	var err error
	_, err = engine.Get(&user)

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
	var user User
	var err error

	if id > 0 {
		user.Id = id
		_, err = engine.Get(&user)
	} else if len(username) > 0 {
		user.Username = username
		_, err = engine.Get(&user)
	} else if len(email) > 0 {
		user.Email = email
		_, err = engine.Get(&user)
	} else if len(nickname) > 0 {
		user.Nickname = nickname
		_, err = engine.Get(&user)
	} else {
		return user, errors.New("at least one parm")
	}

	return user, err
}

// get avatar
func (this *User) GetAvatar(id int, username string, email string, nickname string) (string, error) {
	user, err := this.GetUser(id, username, email, nickname)
	if nil == err {
		addr, _ := global.Config.GetValue("custom", "avatar")
		return addr + com.Md5(user.Email), err
	} else {
		log.WithFields(log.Fields{
			"msg": err,
		}).Warn("GetAvatar Failed.")
		return "", err
	}
}
