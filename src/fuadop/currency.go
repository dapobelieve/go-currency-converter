package fuadop

type Naira float32 // Nigerian Naira

type Dirham float32         // UAD Dirham
type Peso float32           // Argentina Peso
type Real float32           // Brazilian Real
type Pound float32          // GBP
type Euro float32           // Euro
type Yen float32            // Japanese Yen
type CanadianDollar float32 // Canadian Dollar
type Rupee float32          // Indian Rupee
type Nakfa float32          // Eritrean Nakfa
type Florin float32         // Aruban Florin

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00849)
}

func (n Naira) Peso() Peso {
	return Peso(n * 0.34013)
}

func (n Naira) Real() Real {
	return Real(n * 0.01247)
}

func (n Naira) Pound() Pound {
	return Pound(n * 0.00208)
}

func (n Naira) Euro() Euro {
	return Euro(n * 0.00236)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33419)
}

func (n Naira) CanadianDollar() CanadianDollar {
	return CanadianDollar(n * 0.00318)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.18834)
}

func (n Naira) Nakfa() Nakfa {
	return Nakfa(n * 0.03467)
}

func (n Naira) Florin() Florin {
	return Florin(n * 0.00414)
}
