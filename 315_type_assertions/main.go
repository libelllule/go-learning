package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

// Конструкция v, ok := i.(T) называется "comma ok idiom"

/*
Утверждение типа предоставляет доступ к базовому конкретному значению интерфейса.
t := i.(T)
Это выражение утверждает, что значение интерфейса i содержит конкретный тип T, и присваивает базовое значение T переменной t.
Если i не содержит T, это выражение вызовет панику.
*/
