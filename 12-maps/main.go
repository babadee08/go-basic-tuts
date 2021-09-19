package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as keys
	phonebook := map[int]string{
		267586967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267586967])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[845775485] = "yoshi"
	fmt.Println(phonebook)
}
