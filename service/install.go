package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gojudge/ojsite/global"
)

var engine *xorm.Engine

func DBInit() {
	var err error
	dburl, _ := global.Config.GetValue("database", "dburl")
	dbport, _ := global.Config.GetValue("database", "dbport")
	dbuser, _ := global.Config.GetValue("database", "dbuser")
	dbpwd, _ := global.Config.GetValue("database", "dbpwd")
	dbname, _ := global.Config.GetValue("database", "dbname")
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpwd, dburl, dbport, dbname))
	if err != nil {
		fmt.Println(err)
	}
	err = engine.Ping()
	if err != nil {
		fmt.Println("connect error")
		fmt.Println(err)
	}

}

func InstallApp() {

}
