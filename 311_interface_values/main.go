package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()
	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}

//На самом низком уровне Go хранит интерфейс как пару (кортеж) из двух элементов: само значение и ярлык с его точным типом.
//Когда вы вызываете метод через интерфейс (например, i.M()), Go заглядывает в этот кортеж, смотрит на type (тип), находит подходящую реализацию метода именно для этого типа, и запускает её, передавая ей value (значение).
