package output

import "unicode"

func F(str string, maxLen int) string {
	lNow := 0
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			lNow += 2
		} else {
			lNow += 1
		}
	}
	if lNow < maxLen {
		for i := 0; i < (maxLen - lNow); i++ {
			str += " "
		}
	}
	return str
}
