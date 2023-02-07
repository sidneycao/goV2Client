package storage

import (
	"log"
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
	d := OpenFile(configDir)
	if !d.isExist {
		log.Printf("the config dir %s is not exists, creating...\n", configDir)
		err := d.CreateDir()
		if err != nil {
			log.Panic("failed to create the config dir... ")
		}
	}
	// 读取配置文件模板  如果不存在就创建
	f := OpenFile(configModule)
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
	f := OpenFile(configModule)
	err := f.Write(W_NEW, []string{module})
	if err != nil {
		log.Panic("failed to create config module file...")
	}
}

func WriteConfig(d string, fileName string) {
	localFile := configDir + "/" + fileName
	f := OpenFile(localFile)

	if f.isExist {
		f.Delete()
	}

	err := f.Write(W_NEW, []string{d})
	if err != nil {
		log.Panicf("write to config file %s because of %e", localFile, err)
	}
}
