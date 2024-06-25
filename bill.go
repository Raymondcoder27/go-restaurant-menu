package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// formatted string (receiver function)
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v ...$%0.2f\n", k+":", v)
		total += v
	}

	//tip
	fs += fmt.Sprintf("%-10v ...$%0.2f\n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-10v ...$%0.2f\n", "total:", total+b.tip)

	return fs
}

//update tip
func (b *bill) updateTip(tip float64){
	b.tip = tip
}

//add a new item
func (b *bill) addItem(name string,price float64){
	b.items[name] = price
}

//save receiver function
func (b *bill) save(){
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("You saved the file")
}

//save file to a pdf
func (b *bill) saveToPdf(){
	pdf := gofpdf.New("P","mm","A4","")
	pdf.AddPage()
	pdf.SetFont("Arial","I", 13)

	//split the content into lines
	lines := strings.Split(b.format(), "\n")
	for _, line := range lines {
		pdf.Cell(0, 10, line)
		pdf.Ln(10)
	}

	err := pdf.OutputFileAndClose("Bills/"+b.name+".pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println("You saved the file to pdf")
}