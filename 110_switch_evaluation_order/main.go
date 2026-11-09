package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday() // просто число от 0 до 6
	switch time.Saturday {        // берем в качестве значения константу 5
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}
