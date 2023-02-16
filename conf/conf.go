package conf

import (
	"encoding/json"
	"log"

	"github.com/SidneyCao/goV2Client/tools/storage"
)

var subConfigFile = "/usr/local/etc/goV2Config/sub.json"
var nodeConfigFile = "/usr/local/etc/goV2Config/node.json"
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

func SaveDefaultConfig(node VNode) {
	if node.ConfigJson != "" {
		storage.WriteConfig(node.ConfigJson, "/usr/local/etc/v2ray/default.json")
	}
}
