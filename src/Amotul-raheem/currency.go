package Amotul_raheem

type Naira float32 // Naira

type BTC float64 // Bitcoin

type Pounds float32 //Britain Pounds

type Euro float32 //Euro

type SwedishKrona float32 //Swedish Krona

type CanadianDollar float32 //Canadian Dollar

type Dirham float32 //UAE Dihram

type Yen float32 //Japanese Yen

type Rupee float32 //Indian Rupee

type Peso float32 //Argentina Peso

type Cedis float32 //Ghana cedis

func (n Naira) BTC() BTC {
	return BTC(n / 8397293)
}

func (n Naira) Pounds() Pounds {
	return Pounds(n * 0.00236)
}

func (n Naira) Euro() Euro {
	return Euro(n * 0.00236)
}

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00846)
}

func (n Naira) Sedkronas() SwedishKrona {
	return SwedishKrona(n * 0.02583)
}

func (n Naira) CADDollar() CanadianDollar {
	return CanadianDollar(n * 0.00316)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33475)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.19034)
}

func (n Naira) Peso() Peso {
	return Peso(n * 0.34371)
}

func (n Naira) Cedis() Cedis {
	return Cedis(n * 0.02373)
}
