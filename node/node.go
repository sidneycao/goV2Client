package node

import (
	"fmt"
	"goV2Client/conf"
	"goV2Client/tools/args"
	"goV2Client/tools/b64"
	"goV2Client/tools/output"
	"log"
	"strconv"
	"strings"
)

// 节点相关的方法
func ParseArgs(a []string) {
	args.CheckArgsLen(a, 1)
	switch a[0] {
	//添加订阅
	case "--set", "-s":
		args.CheckArgsLen(a, 1)
		SetNode(a[1])
	case "--list", "-l":
		args.CheckArgsLen(a, 1)
		ListServer()
	default:
		ListServer()
	}
}

// 解析vmess链接为 vmess json
// 返回一个VNode结构体 和 config字符串
func ParseNode(vmessLink string) (*conf.VNodeStruct, string) {
	if strings.Index(vmessLink, "vmess://") == 0 {
		vmessJson := b64.B64Decoder(vmessLink[8:])
		log.Println("got vmess json:", vmessJson)
		return conf.Parse2StructAndConf(vmessJson)
	} else {
		return nil, ""
	}
}

func ListServer() {
	fmt.Println("=============================================================================")
	fmt.Println(
		output.F("ID", 5),
		output.F("别名", 30),
		output.F("地址", 40),
		output.F("端口", 10),
		output.F("类型", 5),
	)
	for i, config := range conf.NodeConfigNow.NodeList {
		if i == conf.NodeConfigNow.Id {
			fmt.Println(
				"\033[32m",
				output.F("["+strconv.Itoa(i)+"]", 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 40),
				//Port由float64转为string
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Type, 5),
				"\033[0m",
			)
		} else {
			fmt.Println(
				output.F(" "+strconv.Itoa(i), 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 40),
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Type, 5),
			)
		}
	}
	fmt.Println("=============================================================================")

}

func SetNode(id string) {

}
