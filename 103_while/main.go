package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// без начального условия и условия после цикла (третьего) цикл for превращается в полный аналог цикла while
