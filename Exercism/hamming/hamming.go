package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var n = 0
	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				n++
			}
		}
		return n, nil
	}
	return 0, errors.New("wrong length")
}
