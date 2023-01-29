package sub

import (
	"fmt"
	"goV2Client/tools/params"
	"log"
)

// 订阅相关的方法
func ParseArgs(args []string) {
	params.CheckArgsLen(args, 1)
	switch args[0] {
	case "add":
		params.CheckArgsLen(args, 3)
		if args[1] != "" && args[2] != "" {
			addSub(args[1], args[2])
		}
	case "update":
		params.CheckArgsLen(args, 2)
		fmt.Println("sub update")
	case "del":
		params.CheckArgsLen(args, 2)
		fmt.Println("sub del")
	case "list":
		params.CheckArgsLen(args, 1)
		fmt.Println("sub list")
	default:
		fmt.Println("sub list")
	}
}

//添加订阅
func addSub(name string, url string) {
	log.Println("starting add sub...")
	fmt.Printf("sub add %s %s\n", name, url)
}
