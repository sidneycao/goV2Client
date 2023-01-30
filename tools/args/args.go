package args

import "log"

func CheckArgsLen(a []string, l int) {
	if len(a) < l {
		log.Panicf("args are not enough, now is %d, need %d", len(a), l)
	}
}
