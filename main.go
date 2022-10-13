package main

import (
	"fmt"

	"github.com/hussain4real/go-currency-converter/src/dapobelieve"
	"github.com/hussain4real/go-currency-converter/src/hussain4real"
	"github.com/hussain4real/go-currency-converter/src/pelumichael"
)

func main() {

	currency := dapobelieve.Dollar(2000)
	fmt.Println(currency)

	//testing the conversion with Euro
	aminuConversion := hussain4real.Euros(2)
	fmt.Println(aminuConversion)

	// Test for Pelumichael Conversion with British Pound
	pelumiConversion := pelumichael.GBP(10)
	fmt.Println(pelumiConversion)
}
