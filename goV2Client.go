package main

import (
	"flag"
	"fmt"
	"goV2Client/sub"
	"os"
)

var (
	subUrl = flag.String("sub", "", "订阅地址默认为空")
)

/**
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
**/

func main() {
	args := os.Args
	if len(args) < 2 {
		help()
		os.Exit(0)
	}
	parseArgs(args[1:])
}

func parseArgs(args []string) {
	switch args[0] {
	case "-h", "--help":
		help()
		os.Exit(0)
	case "-sub":
		sub.ParseArgs(args[1:])
	case "-node":
		fmt.Println("node")
	default:
		help()
		os.Exit(0)
	}
}

func help() {
	fmt.Println(
		`订阅管理：
    -sub add {name} {url}
        添加一个订阅，订阅后节点添加到node list
    -sub update {name}
        更新订阅
    -sub del {name}
        删除订阅
    -sub list 
        查看所有订阅
节点管理：
    -node list
        查看所有节点
    -node set {node_id}
        使用该节点
其他:
    -h, --help
        显示此帮助信息
	`)
}
