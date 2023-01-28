package tcPing

import "time"

type Target struct {
	Host string
	IP   string
	Port int

	Counter  int
	Interval time.Duration
	Timeout  time.Duration
}
