package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}

/*
Преобразование типов
Выражение T(v) преобразует значение v в тип T.

Go считает неявное приведение типов огромным источником багов и запрещает любые неявные приведения типов.
*/
