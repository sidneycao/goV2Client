package main

import (
	"fmt"
	"goV2Client/sub"
	"os"
)

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
