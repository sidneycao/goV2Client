package node

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/sidneycao/goV2Client/conf"
	"github.com/sidneycao/goV2Client/tools/args"
	"github.com/sidneycao/goV2Client/tools/b64"
	"github.com/sidneycao/goV2Client/tools/output"
	"github.com/sidneycao/goV2Client/tools/speedtest"
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
	log.Println("speed testing with github.com/sidneycao/tcping...")
	fmt.Println("=======================================================================================================")
	Speedtest()
	fmt.Println(
		output.F("ID", 5),
		output.F("别名", 30),
		output.F("地址", 35),
		output.F("端口", 10),
		output.F("类型", 10),
		output.F("TCP测速", 5),
	)
	for i, config := range conf.NodeConfigNow.NodeList {
		// speed := speedtest.Start(config.Vmess.Add, fmt.Sprint(config.Vmess.Port.(float64)), "2", "1", 1)
		if i == conf.NodeConfigNow.Id {
			fmt.Println(
				"\033[32m",
				output.F("["+strconv.Itoa(i)+"]", 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 35),
				//Port由float64转为string
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Net, 10),
				output.F(config.Speed, 5),
				"\033[0m",
			)
		} else {
			fmt.Println(
				output.F(" "+strconv.Itoa(i), 5),
				output.F(config.Vmess.Ps, 30),
				output.F(config.Vmess.Add, 35),
				output.F(fmt.Sprint(config.Vmess.Port.(float64)), 10),
				output.F(config.Vmess.Net, 10),
				output.F(config.Speed, 5),
			)
		}
	}
	fmt.Println("=======================================================================================================")

}

// 由set触发将node.json中的某条配置写入v2rayConfigFile
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
		output.F("地址", 35),
		output.F("端口", 10),
		output.F("类型", 10),
	)
	fmt.Println(
		output.F(" "+strconv.Itoa(conf.NodeConfigNow.Id), 5),
		output.F(runConfig.Vmess.Ps, 30),
		output.F(runConfig.Vmess.Add, 35),
		output.F(fmt.Sprint(runConfig.Vmess.Port.(float64)), 10),
		output.F(runConfig.Vmess.Net, 10),
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
	/**
	错误，只修改了形参
	for _, config := range conf.NodeConfigNow.NodeList {
		config.Speed = speedtest.Start(config.Vmess.Add, fmt.Sprint(config.Vmess.Port.(float64)), "2", "1", 1)
	}
	**/

	// 将超时时间转为string
	tStr := strconv.Itoa(speedtest.SpeedTestTimeout)

	for i := 0; i < len(conf.NodeConfigNow.NodeList); i++ {
		config := &conf.NodeConfigNow.NodeList[i]
		go speedtest.Start(config.Vmess.Add, fmt.Sprint(config.Vmess.Port.(float64)), tStr, "1", 1, &config.Speed)
	}

	// 休眠两倍于超时时间的时间
	time.Sleep(time.Duration(speedtest.SpeedTestTimeout) * time.Second)
}
