package main

import (
	"fmt"

	"github.com/go-currency-converter/src/unkletayo"
)

func main() {

	currency := unkletayo.USD(1000)
	fmt.Println(currency)

}
