package main

import (
	"fmt"

	"github.com/Christomesh/go-currency-converter/src/philemon"
	"github.com/hussain4real/go-currency-converter/src/dapobelieve"
)

func main() {
	var naira philemon.Naira = 200000.0
	result := naira.ToUnitedStatesDollar()
	fmt.Printf("Successfully converted %.2f Naira(NGN) to %.2f United States Dollar(USD).\n", naira, result)

	currency := dapobelieve.Dollar(2000)
	fmt.Printf("%f", currency)
}
