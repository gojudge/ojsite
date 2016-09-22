package service

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gojudge/ojsite/global"
)

var Engine *xorm.Engine

// DBInit Initial database and check if there is a database
// 初始化数据库，检查数据库用户名，密码是否正确
func DBInit(url, port, user, pwd, name string) error {
	var err error
	_, err = DBGetEngine(url, port, user, pwd, name)
	if err != nil {
		log.WithFields(log.Fields{"msg": err}).Error("连接数据库失败")
		return err
	}
	global.Config.SetValue("database", "dburl", url)
	global.Config.SetValue("database", "dbport", port)
	global.Config.SetValue("database", "dbuser", user)
	global.Config.SetValue("database", "dbpwd", pwd)
	global.Config.SetValue("database", "dbname", name)
	log.Info("连接数据库成功")
	goconfig.SaveConfigFile(global.Config, "conf/conf.ini")

	// check if there is any data in database, depend on if Table user is exsit
	// 数据库中是否有数据，通过判断是否有user表确定
	// if table user is exsit, go back, otherwise import a sql file
	// 如果数据库中有数据，直接返回，否则导入初始化数据
	log.Info("开始导入数据库")
	_, err = Engine.ImportFile("scripts/goj.sql")
	if err != nil {
		log.WithFields(log.Fields{
			"msg": err,
		}).Error("导入数据库失败")
	}
	log.Info("导入数据库完成")
	return nil
}

// DBGetEngine to get a new xorm engine
func DBGetEngine(url, port, user, pwd, name string) (*xorm.Engine, error) {
	var err error
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, url, port, name))
	if err != nil {
		log.WithFields(log.Fields{
			"msg": err,
		}).Error("创建数据库引擎失败")
		return nil, err
	}
	err = Engine.Ping()
	if err != nil {
		log.WithFields(log.Fields{
			"msg": err,
		}).Error("连接数据库失败")
		return nil, err
	}
	return Engine, nil
}

func DBGetEngineWithConfig() (*xorm.Engine, error) {
	url, _ := global.Config.GetValue("database", "dburl")
	port, _ := global.Config.GetValue("database", "dbport")
	user, _ := global.Config.GetValue("database", "dbuser")
	pwd, _ := global.Config.GetValue("database", "dbpwd")
	name, _ := global.Config.GetValue("database", "dbname")

	return DBGetEngine(url, port, user, pwd, name)
}
