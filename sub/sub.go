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
			AddSub(a[1], a[2])
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
		ListSub()
	default:
		ListSub()
	}
}

// 添加订阅
func AddSub(name string, url string) {
	log.Println("starting add sub...")
	if _, value := conf.SubConfigNow[name]; value {
		log.Panic("sub name already exist...")
	}
	sub := conf.VSub{
		Name: name,
		Url:  url,
	}
	log.Println("sub url is:", sub.Url)
	nodeList := GetSub(sub)
	if len(nodeList) == 0 {
		log.Panic("node list is empty...")
	}
	log.Printf("success, got %d node configs...\n", len(nodeList))
	// 保存数据
	conf.NodeConfigNow.NodeList = append(conf.NodeConfigNow.NodeList, nodeList...)
	conf.SubConfigNow[name] = sub

	conf.WriteLocalConfig(conf.SubConfigNow, conf.NodeConfigNow)

}

// 对订阅链接发起get请求 获取返回后的加密文本
// 解密加密文本 获取节点列表
func GetSub(sub conf.VSub) []conf.VNode {
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
	nodeList, err := ParseSub(string(body), sub.Name)
	if err != nil {
		log.Panic(err)
	}
	return nodeList
}

// 解密加密文本 获取节点列表
func ParseSub(res string, subName string) ([]conf.VNode, error) {
	nodeList := make([]conf.VNode, 0)
	subLinks := strings.Split(b64.B64Decoder(res), "\n")
	vmessLinks := make([]string, 0)

	// 过滤一遍数组，去除不符合要求的字符串
	for _, l := range subLinks {
		if strings.Index(l, "vmess://") == 0 {
			vmessLinks = append(vmessLinks, l)
		}
	}
	if len(vmessLinks) == 0 {
		err := errors.New("no vmess link found")
		return nil, err
	}

	// 逐行处理解密，返回node列表
	for _, l := range vmessLinks {
		v, c := node.ParseNode(l)
		nodeList = append(nodeList, conf.VNode{
			SubName:    subName,
			Source:     l,
			Vmess:      *v,
			ConfigJson: c,
		})
	}
	return nodeList, nil
}

func ListSub() {
	fmt.Println("=======================================================")
	fmt.Println("name          ", "url")
	for name, sub := range conf.SubConfigNow {
		fmt.Println(name, "           ", sub.Url)
	}
	fmt.Println("=======================================================")
}
