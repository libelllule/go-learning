package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	pixels := make([]uint8, dx*dy)
	for y := range matrix {
		matrix[y] = pixels[y*dx : (y+1)*dx]
		for x := range matrix[y] {
			matrix[y][x] = uint8((x + y) / 2)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
