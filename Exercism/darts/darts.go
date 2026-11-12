package darts

import "math"

func Score(x, y float64) int {
	rad := math.Sqrt(x*x + y*y)
	switch {
	case rad <= 1:
		return 10
	case rad <= 5:
		return 5
	case rad <= 10:
		return 1
	default:
		return 0
	}
}
