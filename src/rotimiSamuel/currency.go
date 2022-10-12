package rotimisamuel

type Naira float32

// Type Alias for currency conversion

type Guilder float32

type Dram float32

type Lek float32

type NorwegianKroner float32

type CanadianDollar float32

type Pound float32

type Euro float32

type Yen float32

type Peso float32

type Cedis float32

// Methods for Naira to other currencies

func (n Naira) GuilderConverter() Guilder {
	return Guilder(n * 0.00414)
}

func (n Naira) DramConverter() Dram {
	return Dram(n * 0.90119)
}

func (n Naira) LekConverter() Lek {
	return Lek(n * 0.27367)
}

func Kroners(n Naira) NorwegianKroner {
	return NorwegianKroner(n * 0.02512)
}

func CanadianDollars(n Naira) CanadianDollar {
	return CanadianDollar(n * 0.00319)
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

func (n Naira) Peso() Peso {
	return Peso(n * 0.34371)
}

func (n Naira) Cedis() Cedis {
	return Cedis(n * 0.02373)
}
