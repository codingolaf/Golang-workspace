package main

import (
	"fmt"
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

func main() {
	// fmt.Println(rand.Intn(20))
	fmt.Println(add(34, 21))
	// fmt.Println(swap("first", "second"))
	fmt.Println(split(36))
}
