package conf

import (
	"encoding/json"
	"goV2Client/tools/storage"
	"log"
)

var subConfigFile = "sub.json"
var nodeConfigFile = "node.json"

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
