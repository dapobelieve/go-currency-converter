package lolupapi

type Naira float32

type Dollar float32

type Pounds float32

type Euro float32

type SwedishKrona float32

type CanadianDollar float32

type Dirham float32

type Yen float32

type Rupee float32

type Peso float32

type Cedis float32

func (n Naira) Dollar() Dollar {
	return Dollar(n * 0.0023)
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

func SEDKronas(n Naira) SwedishKrona {
	return SwedishKrona(n * 0.02583)
}

func CADDollar(n Naira) CanadianDollar {
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
