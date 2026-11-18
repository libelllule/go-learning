package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var sum int
	for i, ch := range isbn[:9] {
		if !unicode.IsDigit(ch) {
			return false
		}
		sum += int(ch-'0') * (10 - i)
	}
	if isbn[9] == 'X' {
		sum += 10
	} else if unicode.IsDigit(rune(isbn[9])) {
		sum += int(isbn[9] - '0')
	} else {
		return false
	}
	return sum%11 == 0
}
