package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // спецификаор T вместо него подставляется сам тип например bool
	fmt.Printf("Type: %T Value: %v\n", z, z)           // спецификатор v универсальный он сам понимает какой тип даннрых в него передали и форматирует
	// его стандартным образом
}
