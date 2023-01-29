package sub

import (
	"goV2Client/tools/params"
)

// 订阅相关的方法
func ParseArgs(args []string) {
	params.CheckArgsLen(args, 1)
}
