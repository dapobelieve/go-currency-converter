package main

import (
	"fmt"
	"github.com/kelechi-otu/go-currency-converter/src/dapobelieve"
	"github.com/kelechi-otu/go-currency-converter/src/kelechiotu"
)
func main() {


	currency := dapobelieve.Dollar(2000)
	fmt.Println(currency)

	var naira kelechiotu.Naira = 500
	convertedCurrency := naira.GhanaianNewCedis()
    fmt.Printf("%v naira is equivalent to %v Ghanaian New Cedis.", naira, convertedCurrency)
}
