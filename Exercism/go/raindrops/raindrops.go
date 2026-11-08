package raindrops

import "fmt"

var (
	str3 string = "Pling"
	str5 string = "Plang"
	str7 string = "Plong"
)

func Convert(number int) string {
	var strResult string
	if number%3 == 0 {
		strResult += str3
	}
	if number%5 == 0 {
		strResult += str5
	}
	if number%7 == 0 {
		strResult += str7
	}
	if strResult == "" {
		strResult += fmt.Sprintf("%d", number)
	}
	return strResult
}
