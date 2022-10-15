package ghostdev_labs

type (
	Naira   float64 // Nigerian Niara
	Dollars float64 // US Dollars
	Peso    float64 // Argentain Peso
	Cedis   float64 // Ghana Cedis
	Kwacha  float64 // Zambian Kwacha
	Won     float64 // Korean Won
	Pounds  float64 // British Pounds
	Dirham  float64 // UAE Dirham
	Yen     float64 // Japanese Yen
	Dinar   float64 // Algerian Dinar
	Franc   float64 // Rwandan Franc
)

func (n Naira) Dollars() Dollars {
	return Dollars(n * 0.0023)
}

func (n Naira) Peso() Peso {
	return Peso(n * 0.34373)
}

func (n Naira) Cedis() Cedis {
	return Cedis(n * 0.0238)
}

func (n Naira) Kwacha() Kwacha {
	return Kwacha(n * 0.03648)
}

func (n Naira) Won() Won {
	return Won(n * 3.26034)
}

func (n Naira) Pounds() Pounds {
	return Pounds(n * 0.00207)
}

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00846)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33429)
}

func (n Naira) Dinar() Dinar {
	return Dinar(n * 0.32282)
}

func (n Naira) Franc() Franc {
	return Franc(n * 2.40426)
}
