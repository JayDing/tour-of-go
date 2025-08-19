package main

import (
	"fmt"
	"math"
)

func add(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// newton's method for square root
func sqrt(x float64) float64 {
	z := x
	pz := 0.0
	t := math.Pow(10, -15)
	for i := 0; math.Abs(z-pz) > t; i++ {
		pz = z
		z -= (z*z - x) / (2 * z)
		// fmt.Println(i, z, pz, math.Abs(z-pz), t)
	}

	return pz
}

func main() {
	fmt.Printf("%d\n", add(1, 2))

	a, b := swap("hello", "world")
	fmt.Printf("%s, %s\n", a, b)

	c, d := split(17)
	fmt.Printf("%d, %d\n", c, d)

	fmt.Printf("%g, %g\n", sqrt(2), math.Sqrt(2))
}
