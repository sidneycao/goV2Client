package node

import (
	"fmt"
	"goV2Client/conf"
	"goV2Client/tools/args"
	"goV2Client/tools/b64"
	"log"
	"strings"
)

// 节点相关的方法
func ParseArgs(a []string) {
	args.CheckArgsLen(a, 1)
	switch a[0] {
	//添加订阅
	case "set":
		args.CheckArgsLen(a, 1)
		fmt.Println("node set")
	case "list":
		args.CheckArgsLen(a, 1)
		fmt.Println("node list")
	default:
		fmt.Println("node list")
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

/**

func ListServer() {
	fmt.Println("=======================================================")
	fmt.Println(
		putil.F("ID", 4),
		putil.F("别名", 50),
		putil.F("地址", 24),
		putil.F("端口", 10),
		putil.F("类型", 5),
	)
	for i, config := range conf.ServerConfigNow.ServerList {
		if i == conf.ServerConfigNow.Id {
			fmt.Println(
				putil.F("["+strconv.Itoa(i)+"]", 4),
				putil.F(config.Vmess.Ps, 50),
				putil.F(config.Vmess.Add, 24),
				putil.F(config.Vmess.Port, 10),
				putil.F(config.Vmess.Type, 5),
			)
		} else {
			fmt.Println(
				putil.F(" "+strconv.Itoa(i), 4),
				putil.F(config.Vmess.Ps, 50),
				putil.F(config.Vmess.Add, 24),
				putil.F(config.Vmess.Port, 10),
				putil.F(config.Vmess.Type, 5),
			)
		}
	}
	fmt.Println("=======================================================")

}
**/
