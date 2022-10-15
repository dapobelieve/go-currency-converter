package main

import (
	"fmt"

	"github.com/kingsbloc/go-currency-converter/src/dapobelieve"
	"github.com/kingsbloc/go-currency-converter/src/kingsbloc"
)

func main() {
	currency := dapobelieve.Dollar(2000)
	fmt.Println(currency)
	// base Currency is Naira
	amount := 2000
	mycurrency := kingsbloc.Naira(amount)
	fmt.Printf("%d Naira = %f Euros \n", amount, mycurrency.Euro())
	fmt.Printf("%d Naira = %f Pounds \n", amount, mycurrency.Pound())
	fmt.Printf("%d Naira = %f US Dollars \n", amount, mycurrency.UsDollar())
}
