package conf

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/sidneycao/goV2Client/tools/storage"
)

type VNode struct {
	SubName    string
	Vmess      VNodeStruct //解析后的vmess json的结构
	Source     string      //原始vmess链接  vmess://
	ConfigJson string      //通过解析后的vmess json 转换得到的 v2ray config
	Speed      string      //测速后的延迟
}

type VNodeStruct struct {
	Ps   string      `json:"ps"`
	Add  string      `json:"add"`
	Port interface{} `json:"port"` //这里使用interface是为了兼容string和float64的格式
	ID   string      `json:"id"`
	Aid  string      `json:"aid"`
	Net  string      `json:"net"`
	Type string      `json:"type"`
	Host string      `json:"host"`
	Path string      `json:"path"`
	//TLS  string `json:"tls"`
}

type VNodeConfig struct {
	// Id记录的是当前使用的是哪个节点
	Id       int
	NodeList []VNode
}

var NodeConfigNow VNodeConfig = VNodeConfig{-1, []VNode{}}

// 通过vmess json
// 返回VNode结构和v2ray config
func Parse2StructAndConf(vmessJson string) (*VNodeStruct, string) {
	var v VNodeStruct
	err := json.Unmarshal([]byte(vmessJson), &v)
	if err != nil {
		log.Panicf("failed to unmarshall json to vmess struct because of %e...", err)
	}

	//无论port类型是float64还是string 最后都转为int
	if reflect.TypeOf(v.Port).Kind() == reflect.String {
		n, err := strconv.Atoi(v.Port.(string))
		if err != nil {
			log.Panic(err)
		}
		v.Port = n
	} else if reflect.TypeOf(v.Port).Kind() == reflect.Float64 {
		v.Port = int(v.Port.(float64))
	}

	return &v, Parse2Conf(v)
}

// 将VNode解析为v2ray config
func Parse2Conf(v VNodeStruct) string {
	m := storage.LoadConfigModule()
	m = strings.Replace(m, "{Add}", v.Add, 1)
	//Port 由int转为string
	m = strings.Replace(m, "{Port}", strconv.Itoa(v.Port.(int)), 1)
	m = strings.Replace(m, "{ID}", v.ID, 1)
	m = strings.Replace(m, "{Aid}", v.Aid, 1)
	m = strings.Replace(m, "{Net}", v.Net, 1)
	m = strings.Replace(m, "{Host}", v.Host, 1)
	m = strings.Replace(m, "{Path}", v.Path, 1)

	return m
}
