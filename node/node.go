package node

import (
	"goV2Client/conf"
	"goV2Client/tools/b64"
	"log"
	"strings"
)

// 解析vmess链接
// 返回一个vmess结构体 和 config字符串
func ParseNode(vmessLink string) (*conf.VNodeStruct, string) {
	if strings.Index(vmessLink, "vmess://") == 0 {
		vmessJson := b64.B64Decoder(vmessLink[8:])
		log.Println("got vmess json:", vmessJson)
		return conf.Parse2StructAndConf(vmessJson)
	} else {
		return nil, ""
	}
}
