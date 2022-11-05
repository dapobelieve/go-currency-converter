package main

import (
	"fmt"
	"github.com/aosimeon/go-currency-converter/src/aosimeon"
)

func main() {
	// Convert naira to CAD
	currency := aosimeon.ToCAD(2000)

	// Print result
	fmt.Println(currency)
}
