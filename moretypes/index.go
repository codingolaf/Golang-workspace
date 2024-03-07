package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	fmt.Println(v.X)

	p := &v
	(*p).X = 1e3
	fmt.Println(v)

	// array, fixed size
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice, dynamically sized
	var s []int = primes[1:4]
	fmt.Println(s)
}
