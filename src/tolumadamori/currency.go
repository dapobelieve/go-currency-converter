package tolumadamori

//declaring the naira currency type
type Naira float32

//declaring the Dollar currency type
type Dollar float32

//declaring the British Pound currency type
type Pound float32

//declaring the Argentine Peso currency type
type Peso float32

//declaring the Japanese Yen currency type
type Yen float32

//declaring the BTC currency type
type BTC float32

//declaring the Ethereum currency type
type Ethereum float32

//declaring the Yuan currency type
type Yuan float32

//declaring the Euro currency type
type Euro float32

//declaring the Ghanaian Cedi currency type
type Cedi float32

//declaring the Brazilian Real currency type
type Real float32

//This method returns the Naira amount converted to Dollars
func (n Naira) NairaToDollar(amount Naira) Dollar {

	return Dollar(amount * 0.0023)
}

//This method returns the Naira amount converted to British Pounds
func (n Naira) NairaToPound(amount Naira) Pound {

	return Pound(amount * 0.00207)
}

//This method returns the Naira amount converted to Argentine Pesos
func (n Naira) NairaToPeso(amount Naira) Peso {

	return Peso(amount * 0.34373)
}

//This method returns the Naira amount converted to Japanese Yen
func (n Naira) NairaToYen(amount Naira) Yen {

	return Yen(amount * 0.33429)
}

//This method returns the Naira amount converted to Bitcoin
func (n Naira) NairaToBitcoin(amount Naira) BTC {

	return BTC(amount * 0.00000001164299955337454)
}

//This method returns the Naira amount converted to Ethereum
func (n Naira) NairaToEthereum(amount Naira) Ethereum {

	return Ethereum(amount * 0.0000001714239674705879)
}

//This method returns the Naira amount converted to Yuan
func (n Naira) NairaToYuan(amount Naira) Yuan {

	return Yuan(amount * 0.01639)
}

//This method returns the Naira amount converted to Euros
func (n Naira) NairaToEuros(amount Naira) Euro {

	return Euro(amount * 0.00236)
}

//This method returns the Naira amount converted to Ghanaian Cedis
func (n Naira) NairaToCedi(amount Naira) Cedi {

	return Cedi(amount * 0.02411)
}

//This method returns the Naira amount converted to Brazilian Reals
func (n Naira) NairaToReal(amount Naira) Real {

	return Real(amount * 0.01201)
}
