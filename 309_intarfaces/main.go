package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser // переменная a хранит что угодно реализующее интерфейс Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	a = &v
	fmt.Println(a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Сигнатура метода это его описание внешнее без внутреннего устройства
Имя метода (например, Abs).
Входящие параметры: их количество, порядок и типы данных (например, метод ничего не принимает (), или принимает число и строку (n int, s string)).
Возвращаемые значения: их количество и типы (например, возвращает одно число float64 или число и ошибку (int, error)).
*/
