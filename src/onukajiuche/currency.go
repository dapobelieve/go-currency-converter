package onukajiuche

type Naira float32 // Nigerian Naira

type Dirham float32         // UAD Dirham
type Cedi float32           // Ghanaian New Cedi
type Rupee float32           // Seychelles Rupee
type Pound float32          // GBP
type Euro float32           // Euro
type Yen float32            // Japanese Yen
type CanadianDollar float32 // Canadian Dollar
type Shilling float32          // Kenyan Shilling
type Rand float32          // South African Rand
type Won float32         //Korean Won

func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.00849)
}

func (n Naira) Cedi() Cedi {
	return Cedi(n * 0.0245)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.0290)
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

func (n Naira) Shilling() Shilling {
	return Shilling(n * 0.275)
}

func (n Naira) Rand() Rand {
	return Rand(n * 0.0416)
}

func (n Naira) Won() Won {
	return Won(n * 3.280)
}