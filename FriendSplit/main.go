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
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill: ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, t - add tip, f - divide between friends,  s - save calculation): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Input the item name: ", reader)
		price, _ := getInput("Input the item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The prie must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, p)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the bill calculation into a file -  ", b.name)
		fmt.Println("This is the end of your calculation! \n")
	case "t":
		tip, _ := getInput("Input the tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", t)
		promptOptions(b)
	case "f":
		tip, _ := getInput("Did you save the tip y/n? \n", reader)

		if tip == "y" {
			name, _ := getInput("Who gave the money? \n", reader)
			div, _ := getInput("How many friends are there except the provider? \n", reader)
			money, err := strconv.ParseFloat(div, 64)

			if name != "" || err != nil {
				b.provider = name
				divide := b.total / money
				i := 0
				for i < int(money) {
					fname, _ := getInput("Input the friend name: ", reader)
					b.friends[fname] = divide
					i++
				}
			} else {
				promptOptions(b)
			}

		}
		promptOptions(b)

	default:
		fmt.Println("Invalid options.....")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

}
