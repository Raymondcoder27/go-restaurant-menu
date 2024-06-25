package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string,r *bufio.Reader) (string,error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter bill name: ", reader)

	b := newBill(name)
	fmt.Println("Your bill name is", b.name)

	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a - add item, t - add tip, s - save item)", reader)

	switch opt{
	case "a":
		name, _ := getInput("Enter name: ", reader)
		price, _ := getInput("Enter price: ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be of type float")
			promptOptions(b)
		}
		b.addItem(name,p)
		fmt.Println(name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip ($):", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be of type float")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println(tip)
		promptOptions(b)
	case "s":
		b.save()
		b.saveToPdf()
		fmt.Println("You saved the file", b.name)
		// return b
	default:
		fmt.Println("Invalid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	// fmt.Println(mybill)
}