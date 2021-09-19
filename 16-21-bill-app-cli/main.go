package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	// fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOption(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOption(b)
	case "t":
		tip, _ := getInput("Enter tip amount (€): ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOption(b)
		}

		b.updateTip(t)

		fmt.Println("tip added", tip)
		promptOption(b)
	case "s":
		b.save()
		fmt.Println("you save the bill to file", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOption(b)
	}
}

func main() {
	myBill := createBill()
	promptOption(myBill)

	// fmt.Println(myBill)
}
