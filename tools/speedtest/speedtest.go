package speedtest

import (
	"bytes"

	"github.com/SidneyCao/tcping/utils"
)

func Start(host string, port string, timeout string, interval string, counters int) string {
	t := utils.NewTarget(host, port, timeout, interval)
	// 不输出到os.Stdout
	var tmp_buff bytes.Buffer
	pinger := utils.NewPing(*t, &tmp_buff, counters)
	go pinger.Ping()
	// 等待ping完成
	<-pinger.Done()
	// 取整
	return string(utils.Round(pinger.TotalDuration, 1))
}
