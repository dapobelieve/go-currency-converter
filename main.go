package main

import (
	"fmt"
	"github.com/Christomesh/go-currency-converter/src/Christomesh"
)

func main() {

	currency := Christomesh.NGN(9000000)
	fmt.Println(currency.BTC())
}
