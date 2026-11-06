package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//Константы объявляются так же, как переменные, но с использованием ключевого слова const. Константы нельзя объявить с помощью синтаксиса :=
// В отличие от переменных, после создания их невозможно изменить (перезаписать) в процессе работы приложения.
