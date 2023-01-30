package node

import (
	"goV2Client/tools/b64"
	"log"
	"strings"
)

//解析vmess链接
func ParseNode(vmessLink string) string {
	if strings.Index(vmessLink, "vmess://") == 0 {
		vmessJson := b64.B64Decoder(vmessLink[8:])
		log.Println("get vmess json:", vmessJson)
		return vmessJson
	} else {
		return ""
	}
}
