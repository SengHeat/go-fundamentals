package main

import "fmt"

func main() {
	var i int = 23
	var f float64 = 3.14
	var s string = "Gopher"
	var b bool = true
	var by byte = 'A' // alias for uint8 → 65
	var r rune = '界' // alias for int32 (Unicode)

	// Type conversion is always EXPLICIT — no implicit casting
	x := float64(i) + f // must cast int → float64
	y := int(f)         // truncates: 3.14 → 3

	fmt.Println(i, f, s, b, by, r, x, y, r)
}
