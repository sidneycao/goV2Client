package speedtest

import (
	"bytes"
	"fmt"
	"time"

	"github.com/SidneyCao/tcping/utils"
)

var SpeedTestTimeout int = 1

func Start(host string, port string, timeout string, interval string, counters int, result *string) {
	t := utils.NewTarget(host, port, timeout, interval)
	// 不输出到os.Stdout
	var tmp_buff bytes.Buffer
	pinger := utils.NewPing(*t, &tmp_buff, counters)
	pinger.Ping()
	// 取整
	// return fmt.Sprint(utils.Round(pinger.TotalDuration, 1))
	// 如果时间大于等于SpeedTestTimeout，或者失败次数等于ping个数，就认定为失败
	if pinger.TotalDuration >= time.Duration(SpeedTestTimeout)*time.Second || pinger.Failed == counters {
		*result = "Failed"
	} else {
		*result = fmt.Sprint(utils.Round(pinger.TotalDuration, 1))
	}
}
