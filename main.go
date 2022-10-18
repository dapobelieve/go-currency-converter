package main

import (
	"fmt"
	"github.com/kelechi-otu/go-currency-converter/src/dapobelieve"
	"github.com/kelechi-otu/go-currency-converter/src/kelechiotu"
)
func main() {


	currency := dapobelieve.Dollar(2000)
	fmt.Println(currency)

	var sampleNaira kelechiotu.Naira = 500
	convertedCurrency := sampleNaira.GhanaianNewCedis()
    fmt.Println(convertedCurrency)
}
