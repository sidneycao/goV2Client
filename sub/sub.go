package sub

import (
	"encoding/base64"
	"errors"
	"fmt"
	"goV2Client/conf"
	"goV2Client/tools/params"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// 订阅相关的方法
func ParseArgs(args []string) {
	params.CheckArgsLen(args, 1)
	switch args[0] {
	//添加订阅
	case "add":
		params.CheckArgsLen(args, 3)
		if args[1] != "" && args[2] != "" {
			addSub(args[1], args[2])
		} else {
			log.Panicf("args error...")
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

//解析base64
func parseSub(res string, subName string) ([]conf.V2Node, error) {
	nodeList := make([]conf.V2Node, 0)
	b64 := make([]byte, base64.RawStdEncoding.DecodedLen(len(res)))
	d, err := base64.StdEncoding.Decode(b64, []byte(res))
	if err != nil {
		log.Panic(err)
	}
	subLinks := strings.Split(string(b64[:d]), "\n")
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
		nodeList = append(nodeList, conf.V2Node{
			SubName: subName,
			Source:  l,
			//ConfigJson: "123",
		})
	}
	return nodeList, nil
}
