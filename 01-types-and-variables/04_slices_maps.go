package main

import "fmt"

func main() {
	// Slice - dynamic array(pointer + len + cap)
	slice := []string{"go", "is", "fun"}
	slice = append(slice, "and", "fast")

	// make — pre-allocate capacity for performance
	numbers := make([]int, 0, 10)
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i*i)
	}

	scores := map[string]int{
		"alice": 79,
		"bob":   90,
	}
	scores["charlie"] = 85
	delete(scores, "bob")

	val, exists := scores["bob"]
	if !exists {
		fmt.Println("dave not found, zero val:", val)
	}

	fmt.Println(slice, numbers, scores)
}
