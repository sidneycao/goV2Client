package storage

import (
	"log"
	"os"
)

const configDir = "/usr/local/etc/goV2Config"

var configModule string = configDir + "/" + "config_module.json"

var module = `{ 
	"outbounds": [
		{
			"protocol": "vmess",
			"tag": "default",
			"settings": {
				"vnext": [
					{
						"address": "{Add}",
						"port": {Port},
						"users": [
							{
								"encryption": "none",
								"id": "{ID}",
								"alterId": {Aid},
								"security": "auto"
							}
						]
					}
				]
			},
			"streamSettings":{
				"sockopt": {
					"mark": 255
				}
			}
		}
	]
}`

// 加载配置文件模板
func LoadConfigModule() string {
	// 检查配置文件目录是否存在
	d := Open(configDir)
	if !d.isExist {
		log.Printf("the config dir %s is not exists, creating...\n", configDir)
		err := os.MkdirAll(d.path, os.ModePerm)
		if err != nil {
			log.Panic("failed to create the config dir... ")
		}
	}
	// 读取配置文件模板  如果不存在就创建
	f := Open(configModule)
	r, err := f.Read()
	if err == nil && string(r) != "" {
		return string(r)
	} else {
		CreateConfigModule()
		return module
	}
}

// 创建配置文件模板
func CreateConfigModule() {
	f := Open(configModule)
	err := f.Write(W_NEW, []string{module})
	if err != nil {
		log.Panic("failed to create config module file...")
	}
}

func WriteConfig(d string, path string) {
	f := Open(path)
	if !f.isExist {

	}

}
