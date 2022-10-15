package supe-code

type Naira float32

type Franc float32
type Yen float32
type Baht float32
type Rand float32
type Riyal float32
type Euro float32
type Rupee float32
type Ringgit float32
type Krona float32
type Dirham float32

func (n Naira) Franc() Franc {
	return Franc(n * 0.0023)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33419)
}

func (n Naira) Baht() Baht {
	return Baht(n * 0.088)
}

func (n Naira) Rand() Rand {
	return Rand(n * 0.042)
}

func (n Naira) Riyal() Riyal {
	return Riyal(n * 0.0086)
}

func (n Naira) Euro() Euro {
	return Euro(n * 0.00236)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.18834)
}

func (n Naira) Ringgit() Ringgit {
	return Ringgit(n * 0.011)
}

func (n Naira) Krona() Krona {
	return Krona(n * 0.026)
}

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00849)
}
