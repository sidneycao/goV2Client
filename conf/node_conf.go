package conf

import (
	"encoding/json"
	"fmt"
	"goV2Client/tools/storage"
	"log"
	"strconv"
	"strings"
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
func ParseVmess2StructConf(vmessJson string) (*VmessStuct, string) {
	var v VmessStuct
	err := json.Unmarshal([]byte(vmessJson), &v)
	if err != nil {
		log.Panic("failed to unmarshall json to vmess struct...")
	}
	fmt.Println("success unmarshall")
	return &v, ParseVmess2Conf(v)
}

func ParseVmess2Conf(v VmessStuct) string {
	m := storage.LoadConfigModule()
	m = strings.Replace(m, "{Add}", v.Add, 1)
	m = strings.Replace(m, "{Port}", strconv.Itoa(v.Port), 1)
	m = strings.Replace(m, "{ID}", v.ID, 1)
	m = strings.Replace(m, "{Aid}", v.Aid, 1)

	return m
}
