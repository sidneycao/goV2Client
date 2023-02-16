package node

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/SidneyCao/goV2Client/conf"
	"github.com/SidneyCao/goV2Client/tools/args"
	"github.com/SidneyCao/goV2Client/tools/b64"
	"github.com/SidneyCao/goV2Client/tools/output"
	"github.com/SidneyCao/goV2Client/tools/speedtest"
)

// 节点相关的方法
func ParseArgs(a []string) {
	args.CheckArgsLen(a, 1)
	switch a[0] {
	//添加订阅
	case "--set", "-s":
		args.CheckArgsLen(a, 1)
		SetNode(a[1])
	case "--list", "-l":
		args.CheckArgsLen(a, 1)
		ListServer()
	default:
		ListServer()
	}
}

// 解析vmess链接为 vmess json
// 返回一个VNode结构体 和 config字符串
func ParseNode(vmessLink string) (*conf.VNodeStruct, string) {
	if strings.Index(vmessLink, "vmess://") == 0 {
		vmessJson := b64.B64Decoder(vmessLink[8:])
		log.Println("got vmess json:", vmessJson)
		return conf.Parse2StructAndConf(vmessJson)
	} else {
		return nil, ""
	}
}

func ListServer() {
	fmt.Println("=============================================================================")
	Speedtest()
	fmt.Println(
		output.F("ID", 5),
		output.F("别名", 30),
		output.F("地址", 40),
		output.F("端口", 10),
		output.F("类型", 5),
		output.F("TCP测速", 5),
	)
	for i, config := range conf.NodeConfigNow.NodeList {
		// 此处方法不合理 待优化
		// speed := speedtest.Start(config.Vmess.Add, fmt.Sprint(config.Vmess.Port.(float64)), "2", "1", 1)
		if i == conf.NodeConfigNow.Id {
			fmt.Println(
				"\033[32m",
				output.F("["+strconv.Itoa(i)+"]", 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 40),
				//Port由float64转为string
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Net, 5),
				output.F(config.Speed, 5),
				"\033[0m",
			)
		} else {
			fmt.Println(
				output.F(" "+strconv.Itoa(i), 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 40),
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Net, 5),
				output.F(config.Speed, 5),
			)
		}
	}
	fmt.Println("=============================================================================")

}

// 由set触发将node.json中的某条配置写入default.json
func SetNode(id string) {
	i, _ := strconv.Atoi(id)
	if i < 0 || i >= len(conf.NodeConfigNow.NodeList) {
		log.Panicf("node id [%d] err...", i)
	}
	conf.NodeConfigNow.Id = i
	conf.WriteLocalConfig(conf.SubConfigNow, conf.NodeConfigNow)
	conf.SaveDefaultConfig(conf.NodeConfigNow.NodeList[i])
	log.Printf("success to set node id [%d]...\n", i)
	RestartV2ray()
}

func RestartV2ray() {
	log.Println("restarting v2ray process...")
	runConfig := conf.NodeConfigNow.NodeList[conf.NodeConfigNow.Id]

	log.Println("now config is:  ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓ ↓")
	fmt.Println("=============================================================================")
	fmt.Println(
		output.F("ID", 5),
		output.F("别名", 30),
		output.F("地址", 40),
		output.F("端口", 10),
		output.F("类型", 5),
	)
	fmt.Println(
		output.F(" "+strconv.Itoa(conf.NodeConfigNow.Id), 5),
		output.F(runConfig.Vmess.Ps, 30),
		output.F(runConfig.Vmess.Add, 40),
		output.F(fmt.Sprint(runConfig.Vmess.Port.(float64)), 10),
		output.F(runConfig.Vmess.Net, 5),
	)
	fmt.Println("=============================================================================")

	cmd := exec.Command("systemctl", "restart", "v2ray")
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
	log.Println("success to restart v2ray process")
}

func Speedtest() {
	for _, config := range conf.NodeConfigNow.NodeList {
		config.Speed = speedtest.Start(config.Vmess.Add, fmt.Sprint(config.Vmess.Port.(float64)), "2", "1", 1)
	}
}
