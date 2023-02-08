package main

import (
	"fmt"
	"goV2Client/conf"
	"goV2Client/node"
	"goV2Client/sub"
	"os"
)

func main() {
	a := os.Args
	if len(a) < 2 {
		help()
		os.Exit(0)
	}
	//初始化读取数据
	conf.LoadLocalConfig()
	parseArgs(a[1:])
}

func parseArgs(a []string) {
	switch a[0] {
	case "-h", "--help":
		help()
		os.Exit(0)
	case "--sub":
		sub.ParseArgs(a[1:])
	case "--node":
		node.ParseArgs(a[1:])
	default:
		help()
		os.Exit(0)
	}
}

func help() {
	fmt.Println(
		`订阅管理：
    --sub --add(-a) {name} {url}
        添加一个订阅，订阅后节点添加到node list
    --sub --update(-u) {name}
        更新订阅
    --sub --del(-d) {name}
        删除订阅
    --sub --list(-l) 
        查看所有订阅
节点管理：
    --node --list(-l)
        查看所有节点
    --node --set(-s) {node_id}
        使用该节点
其他:
    -h, --help
        显示此帮助信息
	`)
}
