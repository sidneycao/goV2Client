package speedtest

import (
	"time"

	"github.com/SidneyCao/tcping/utils"
)

func Start(host string, port string, timeout string, interval string, counters int) time.Duration {
	t := utils.NewTarget(host, port, timeout, interval)
	pinger := utils.NewPing(*t, counters)
	go pinger.Ping()
	// 等待ping完成
	<-pinger.Done()
	// 取整
	return utils.Round(pinger.TotalDuration, 0)
}
