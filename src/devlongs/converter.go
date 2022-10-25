package devlongs

type Naira   float64 
type Pounds  float64 
type Yen     float64 
type Dinar   float64
type Franc   float64 
type Dollars float64 
type Won     float64 
type Peso    float64 
type Cedis   float64 
type Kwacha  float64 
type Dirham  float64 

// conversion to dollars
func (n Naira) Dollars() Dollars {
	return Dollars(n * 0.0023)
}

// conversion to Peso
func (n Naira) Peso() Peso {
	return Peso(n * 0.34373)
}

// conversion to Cedis
func (n Naira) Cedis() Cedis {
	return Cedis(n * 0.0238)
}

// conversion to Kwacha
func (n Naira) Kwacha() Kwacha {
	return Kwacha(n * 0.03648)
}

// conversion to Won
func (n Naira) Won() Won {
	return Won(n * 3.26034)
}

// conversion to Pounds
func (n Naira) Pounds() Pounds {
	return Pounds(n * 0.00207)
}

// conversion to Dirham
func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00846)
}

// conversion to Yen
func (n Naira) Yen() Yen {
	return Yen(n * 0.33429)
}

// conversion to Dinar
func (n Naira) Dinar() Dinar {
	return Dinar(n * 0.32282)
}

// conversion to Franc
func (n Naira) Franc() Franc {
	return Franc(n * 2.40426)
}
