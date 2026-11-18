package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	var sum int
	isSecond := len(id)%2 == 0
	for _, ch := range id {
		if !unicode.IsDigit(ch) {
			return false
		}
		val := int(ch - '0')
		if isSecond {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		isSecond = !isSecond
		sum += val
	}
	return sum%10 == 0
}
