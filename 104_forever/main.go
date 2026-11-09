package main

import "fmt"

func main() {
	var counter = 0
	fmt.Println("Start cycle")
	for {
		counter += 1
		if counter == 3 {
			fmt.Println("Equal 3")
			break
		}
	}
	fmt.Println("Finish cycle")
}

// цикл вообще без условий это бесконечный цикл
