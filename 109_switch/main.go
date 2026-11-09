package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on")
	// можн как в цикле for написать начальное условие
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

}

// В отличие от С в Go в конструкции switch case могут проверяться не только целые числа и символы но и целые строки и числа с плавающей точкой
// Кроме того не надо прописывать break отдельно ( в отличие от С) цикл сам прервется после того как будет выполнен один case
