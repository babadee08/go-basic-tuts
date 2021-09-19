package main

import "fmt"

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	// fmt.Println("Hello, world!")
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint8 = 255

	// var scoreOne float32 = 25.96
	// var scoreTwo float64 = 234567887654345676543345.7
	// scoreThree := 1.5

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New line \n")

	age := 35
	name := "shaun"

	// Println
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (Formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)

}
