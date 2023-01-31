package conf

import (
	"encoding/json"
	"log"
)

type V2Node struct {
	SubName    string
	Vmess      VmessStuct //解析后的vmess json的结构
	Source     string     //原始vmess链接  vmess://
	ConfigJson string     //通过解析后的vmess json 转换得到的 v2ray config
}

type VmessStuct struct {
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port int    `json:"port"`
	ID   string `json:"id"`
	Aid  string `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	//TLS  string `json:"tls"`
}

// 通过vmess json
// 返回vmess结构 和 v2ray config
func ParseVmess2StructConf(vmessJson string) *VmessStuct {
	var v VmessStuct
	err := json.Unmarshal([]byte(vmessJson), &v)
	if err != nil {
		log.Panic(err)
	}
	return &v
}

func ParseVmess2Conf() {

}
