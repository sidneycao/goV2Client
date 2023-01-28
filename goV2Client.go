package main

import (
	"bufio"
	"flag"
	"fmt"
	"goV2Client/b64Decoder"
	"goV2Client/curl"
	"strings"
)

var (
	subUrl = flag.String("sub", "", "订阅地址默认为空")
)

func main() {
	//解析命令行参数
	flag.Parse()

	//通过订阅url获取base64加密后的返回
	res := curl.Cget(*subUrl)

	//解密base64  获取节点
	nodes := b64Decoder.Decoder([]byte(res))

	//逐行解析出节点
	scanner := bufio.NewScanner(strings.NewReader(string(nodes)))
	for scanner.Scan() {
		decodeStr := strings.Split(scanner.Text(), "//")[1]
		node := b64Decoder.Decoder([]byte(decodeStr))
		fmt.Println(string(node))
	}
}
