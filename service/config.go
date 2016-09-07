package service

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func ConfigLoad() (*goconfig.ConfigFile, error) {
	fmt.Println("configLoad")
	configFile, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return configFile, nil
}
