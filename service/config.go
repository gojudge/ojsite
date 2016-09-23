package service

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gojudge/ojsite/global"
	"os"
)

func ConfigLoad() (*goconfig.ConfigFile, error) {
	fmt.Println("configLoad")
	filename := "conf/conf.ini"
	fin, err := os.Open(filename)
	defer fin.Close()
	if err != nil {
		fout, err := os.Create(filename)
		defer fout.Close()
		fmt.Println(err)
	}
	configFile, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return configFile, nil
}

func ConfigIsInit() bool {
	dburl, _ := global.Config.GetValue("database", "dburl")
	return dburl != ""
}
