package main

import (
	"fmt"

	"github.com/UnkleTayo/go-currency-converter/src/unkletayo"
)

func main() {
	currency := unkletayo.Dollar(2000)
	fmt.Println(currency)
}
