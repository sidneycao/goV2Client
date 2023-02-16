package speedtest

import (
	"time"

	"github.com/SidneyCao/tcping/utils"
)

func Start(host string, port string, timeout string, interval string, counters int) time.Duration {
	t := utils.NewTarget(host, port, timeout, interval)
	pinger := utils.NewPing(*t, counters)
	go pinger.Ping()
	return pinger.TotalDuration
}
