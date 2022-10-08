package main

import (
	"./src/kingsbloc"
)

func main() {
	// base Currency is Naira
	currency := kingsbloc.Naira(2000)
	currency.Euro()
}
