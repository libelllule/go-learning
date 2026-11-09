package main

import "fmt"

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// вызовы defer кладутся в стек выполняются по принципу LIFO
