package main

import "fmt"

func main() {

	// Classic for loop
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//While style
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)

	// Range over slice — index and value
	items := []string{"Apple", "Banana", "Orange"}
	for i, t := range items {
		fmt.Printf("%d: %s\n", i, t)
	}

	// Range over map — order is RANDOM every run
	m := map[string]string{
		"name":    "Seng Heat",
		"country": "Cambodia",
	}

	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
