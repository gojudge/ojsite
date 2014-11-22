package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gogather/com"
)

// fis map
func Fis(key string) string {
	content := loadMap()
	json, _ := com.JsonDecode(content)
	json = json.(map[string]interface{})["res"]
	if fileMap, ok := json.(map[string]interface{}); !ok {
		fmt.Println("map.json id illeage!")
	} else {
		for tmpKey, views := range fileMap {
			uri, ok := views.(map[string]interface{})["uri"].(string)
			if !ok {
				fmt.Println("error in map.json")
			}

			fileType, ok := views.(map[string]interface{})["type"].(string)
			if !ok {
				fmt.Println("error in map.json")
			}

			if tmpKey == key {
				if fileType == "css" {
					return `<link rel="stylesheet" href="/static` + uri + `">`
				} else if fileType == "js" {
					return `<script src="/static` + uri + `"></script>`
				}
			}
		}
	}

	return ""
}

// load map.json
func loadMap() string {
	mapPath := beego.AppConfig.String("static_map")
	mapContent := com.ReadFile(mapPath)
	return mapContent
}
