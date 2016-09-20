package service

import (
	"fmt"
	"github.com/Unknwon/goconfig"
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
