package sunday4me

// naira is converted with the below currency with the rate of 1naira to the below currency
// oct 4th 2022

type Naira float32 // Nigerian Naira

type Pound float32          // GBP
type Euro float32           // Euro
type Rupee float32          // Indian Rupee
type BarbadosDollar float32 // Indian Rupee
type CubanPeso float32      // Cuban Peso
type Dirham float32         // UAD Dirham
type Peso float32           // Argentina Peso
type Real float32           // Brazilian Real
type Yen float32            // Japanese Yen
type CanadianDollar float32 // Canadian Dollar

func (n Naira) Pound() Pound {
	return Pound(n * 0.00208)
}

func (n Naira) Euro() Euro {
	return Euro(n * 0.00236)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.18834)
}

func (n Naira) BarbadosDollar() BarbadosDollar {
	return BarbadosDollar(n * 0.00462)
}

func (n Naira) CubanPeso() CubanPeso {
	return CubanPeso(n * 0.05776)
}

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00849)
}

func (n Naira) Peso() Peso {
	return Peso(n * 0.34013)
}

func (n Naira) Real() Real {
	return Real(n * 0.01247)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33419)
}

func (n Naira) CanadianDollar() CanadianDollar {
	return CanadianDollar(n * 0.00318)
}

//comment on try and errors of the code below

// 1. conversion of Naira to Dollar
// type Naira float32

// type Dollar float32

///func (n Naira) Dollars() Dollar {
// return Dollar(n * 650)
// }

// 2. conversion of Yens to Pounds
// type Yen float32

// type Pound float32

// func (y Yen) Pounds() Pound {
// 	return Pound(y * 160)
//}
// type Florin float32         // Aruban Florin

//func (n Naira) Florin() Florin {
//return Florin(n * 0.00414)
// }

// trying to fix the code 