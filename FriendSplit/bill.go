package main

import (
	"fmt"
	"io/ioutil"
)

type bill struct {
	name     string
	provider string
	item     map[string]float64
	friends  map[string]float64
	tip      float64
	total    float64
}

// make new bills
func newBill(name string) bill {
	b := bill{}
	b.name = name
	b.provider = "root"
	b.item = map[string]float64{}
	b.friends = map[string]float64{}
	b.total = 0
	b.tip = 0
	return b
}

func (b *bill) format() string {
	fs := "Bill Breakdown: \n"

	for k, v := range b.item {
		fs += fmt.Sprintf("%-25v .... $%0.2f \n", k+":", v)
	}
	fs += fmt.Sprintf("%-25v .... $%0.2f \n", "Tip:", b.tip)
	fs += fmt.Sprintf("%-25v .... $%0.2f \n", "Total:", b.total)

	fs += fmt.Sprint("------------------------------------------ \n")
	fs += fmt.Sprintf("Friends who owes money to %v who gave %0.2f \n", b.provider, b.total)
	for k, v := range b.friends {
		fs += fmt.Sprintf("%-25v .... $%0.2f \n", k+":", v)
	}
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
	b.total += b.tip
}

func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
	b.total += price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := ioutil.WriteFile(b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("The bill was saved to file!")
}
