package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// shortened
//
//	func add(a, b int) int {
//		return a + b
//	}
func add(a int, b int) int {
	return a + b
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// variables declared by the statement are only in scope until the end of the `if`.
// variables are also available inside any of the `else` blocks.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// newton's method for square root
func sqrt(x float64) float64 {
	z := x / 2
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
	// fmt.Println(math.pi) // non exported
	fmt.Println(math.Pi) // exported

	fmt.Println(add(1, 2))

	a, b := swap("hello", "world")
	fmt.Printf("%s, %s\n", a, b)

	c, d := split(17)
	fmt.Printf("%d, %d\n", c, d)

	// variable declaration
	var i int = 1
	// shortened variable declaration
	j := 1
	fmt.Println(i, j)

	// implicit initial value
	var ii int     // 0
	var ff float64 // 0.0
	var bb bool    // false
	var ss string  // ""
	fmt.Printf("%v, %v, %v, %v\n", ii, ff, bb, ss)

	// type inference
	x := 42           // int
	y := 3.142        // float64
	z := 0.867 + 0.5i // complex128
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// constants, cannot be declared with :=
	const Constant string = "constant"
	fmt.Println("Happy", Constant, "Day")

	// an untyped constant takes the type needed by its context.
	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Printf("%g, %g\n", sqrt(2), math.Sqrt(2))

	// no break statement in Go
	// variables declared by the statement are only in scope until the end of the `switch` or `case`.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch with no condition is the same as `switch true`
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
