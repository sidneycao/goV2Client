package params

import "log"

func CheckArgsLen(args []string, l int) {
	if len(args) < l {
		log.Panicf("args are not enough, now is %d, need %d", len(args), l)
	}
}
