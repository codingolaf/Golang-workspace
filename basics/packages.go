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

func add(a, b int) int {
	return a + b
}

// multiple results
func swap(a, b string) (string, string) {
	return b, a
}

// named results
func split(sum int) (a, b int) {
	a = sum / 4
	b = sum - a
	return
}

// type conversion
func converter(a int) (f float64, u uint) {
	f = float64(a)
	u = uint(a)
	return
}

func main() {
	// fmt.Println(rand.Intn(20))
	// fmt.Println(add(34, 21))
	// fmt.Println(swap("first", "second"))
	// fmt.Println(split(36))
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(converter(-12))
}
