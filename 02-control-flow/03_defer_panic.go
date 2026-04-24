package main

import "fmt"

func processFile(name string) {
	fmt.Println("opening:", name)
	defer fmt.Println("closing:", name) // runs last
	defer fmt.Println("cleanup B")      // runs second
	defer fmt.Println("cleanup A")      // runs first (LIFO)
	fmt.Println("doing work on:", name)
}

// recover must be inside a deferred function
func safeRun(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic caught: %v", r)
		}
	}()

	fn()

	return nil
}

func main() {
	processFile("data.csv")

	err := safeRun(func() {
		panic("something broke")
	})
	fmt.Println("handled:", err)
}
