package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}

// Тип *T — это указатель на значение типа T. Его нулевое значение (zero value) — nil.
// var p *int

// Оператор * обозначает базовое значение указателя (значение, на которое он указывает).
// Это известно как "разыменование" (dereferencing) или "косвенное обращение" (indirecting).
//В отличие от C, в Go нет адресной арифметики (pointer arithmetic).
