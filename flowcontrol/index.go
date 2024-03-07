package main

import (
	"fmt"
	"math"
	"time"
)

func sumUsingFor() int {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func sumUsingContinued() int {
	sum := 1
	for sum < 1<<20 {
		sum += sum
	}
	return sum
}

func sumForLikeWhile() int {
	sum := 1
	for sum < 1<<20 {
		sum += sum
	}
	return sum
}

func shortIf(a, b, c float64) float64 {
	if v := math.Pow(a, b); v > c {
		return v
	}
	return c
}

func main() {
	t := time.Now()
	defer fmt.Println("I am defered")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 20:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}
