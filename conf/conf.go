package conf

import (
	"encoding/json"
	"goV2Client/tools/storage"
	"log"
)

var subConfigFile = "sub.json"
var nodeConfigFile = "node.json"
var initFlag = false

// 配置文件相关的方法

func WriteLocalConfig(subConfig map[string]VSub, nodeConfig VNodeConfig) {
	subConfigJson, err := json.MarshalIndent(subConfig, "", "    ")
	if err != nil {
		log.Panic("sub config marshall fail...")
	} else {
		storage.WriteConfig(string(subConfigJson), subConfigFile)
	}
	nodeConfigJson, err := json.MarshalIndent(nodeConfig, "", "    ")
	if err != nil {
		log.Panic("sub config marshall fail...")
	} else {
		storage.WriteConfig(string(nodeConfigJson), nodeConfigFile)
	}
}

func LoadLocalConfig() {
	if !initFlag {
		subConfigBytes := storage.ReadConfig(subConfigFile)
		if string(subConfigBytes) != "" {
			err := json.Unmarshal(subConfigBytes, &SubConfigNow)
			if err != nil {
				log.Panicf("failed to unmarshal config %s because of %e", subConfigFile, err)
			}
		}
		nodeConfigBytes := storage.ReadConfig(nodeConfigFile)
		if string(nodeConfigBytes) != "" {
			err := json.Unmarshal(nodeConfigBytes, &NodeConfigNow)
			if err != nil {
				log.Panicf("failed to unmarshal config %s because of %e", nodeConfigFile, err)
			}
		}
	}
	initFlag = true
	log.Println("load config success...")
}
