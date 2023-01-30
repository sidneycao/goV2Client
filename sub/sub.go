package sub

import (
	"errors"
	"fmt"
	"goV2Client/conf"
	"goV2Client/node"
	"goV2Client/tools/args"
	"goV2Client/tools/b64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 订阅相关的方法
func ParseArgs(a []string) {
	args.CheckArgsLen(a, 1)
	switch a[0] {
	//添加订阅
	case "add":
		args.CheckArgsLen(a, 3)
		if a[1] != "" && a[2] != "" {
			addSub(a[1], a[2])
		} else {
			log.Panicf("args error...")
		}
	case "update":
		args.CheckArgsLen(a, 2)
		fmt.Println("sub update")
	case "del":
		args.CheckArgsLen(a, 2)
		fmt.Println("sub del")
	case "list":
		args.CheckArgsLen(a, 1)
		fmt.Println("sub list")
	default:
		fmt.Println("sub list")
	}
}

//添加订阅
func addSub(name string, url string) {
	log.Println("starting add sub...")
	if _, value := conf.SubConfig[name]; value {
		log.Println("sub name already exist...")
		os.Exit(0)
	}
	sub := conf.V2Sub{
		Name: name,
		Url:  url,
	}
	log.Println("sub url is:", sub.Url)
	nodeList := getSub(sub)
	fmt.Println(nodeList)
}

func getSub(sub conf.V2Sub) []conf.V2Node {
	//对订阅链接发起get请求
	req, _ := http.NewRequest("GET", sub.Url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	nodeList, err := parseSub(string(body), sub.Name)
	if err != nil {
		log.Panic(err.Error())
	}
	return nodeList
}

//解析订阅链接
func parseSub(res string, subName string) ([]conf.V2Node, error) {
	nodeList := make([]conf.V2Node, 0)
	subLinks := strings.Split(b64.B64Decoder(res), "\n")
	vmessLinks := make([]string, 0)
	for _, l := range subLinks {
		if strings.Index(l, "vmess://") == 0 {
			vmessLinks = append(vmessLinks, l)
		}
	}
	if len(vmessLinks) == 0 {
		err := errors.New("no vmess link found")
		return nil, err
	}
	for _, l := range vmessLinks {
		config := node.ParseNode(l)
		nodeList = append(nodeList, conf.V2Node{
			SubName:    subName,
			Source:     l,
			ConfigJson: config,
		})
	}
	return nodeList, nil
}
