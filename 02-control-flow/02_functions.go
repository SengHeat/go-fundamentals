package main

import (
	"fmt"
)

// func + funcName + (parameter...) (return)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic — accepts zero or more ints
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	return total
}

// Closure — captures surrounding scope
func makeCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	result, err := divide(10, 5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%.4f\n", result)

	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...)) // spread slice

	c := makeCounter()
	fmt.Println(c(), c(), c()) // 1 2 3
}
