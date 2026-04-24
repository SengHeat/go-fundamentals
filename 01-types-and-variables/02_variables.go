package main

import "fmt"

var globalName = "Go" // package-level, inferred type

func main() {

	// Way 1: explicit type
	var age int = 23

	// Way 2: inferred type
	var name = "Seng Heat"

	// Way 3: short declaration (most common inside functions)
	score := 95.9

	// Zero values — Go always initializes, never garbage
	var count int
	var label string
	var active bool

	fmt.Println(age, name, label, score, count, active)
}
