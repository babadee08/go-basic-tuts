package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// Non-Pointer variables
	// group A types -> ints, strings, bools, floats, arrays and structs
	name := "tifa"

	name = updateName(name)

	fmt.Println(name)

	// Pointer wrapped variables
	// group B types -> slices, maps and functions
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
