package unkletayo

// create a generic currency type
type currency float32

// Declare the currency and match their types to the currency type
type Naira currency

type USD currency             // US Dollar
type SwedishKrona currency    // Swedish Krona
type NorwegianKroner currency // Norwegian Krone
type CanadianDollar currency  // Canadian Dollar
type BritishPound currency    // Great Britain Pound
type Rupee currency           // Indian Rupee
type CongoleseFranc currency  // Congolese Franc
type AngolanKwanza currency   // Angolan Kwanza
type Euro currency            // Euro
type BarbadosDollar currency  // Barbados Dollar

//  receiver functions that converts a naira to another currency

func (n Naira) Dollar() USD {
	return USD(n * 0.00228)
}

func (n Naira) SwedishKrona() SwedishKrona {
	return SwedishKrona(n * 0.02501)
}

func (n Naira) NorwegianKroner() NorwegianKroner {
	return NorwegianKroner(n * 0.02352)
}

func (n Naira) CanadianDollar() CanadianDollar {
	return CanadianDollar(n * 0.0031)
}

func (n Naira) BritishPound() BritishPound {
	return BritishPound(n * 0.00197)
}

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.1879)
}

func (n Naira) CongoleseFranc() CongoleseFranc {
	return CongoleseFranc(n * 4.58114)
}

func (n Naira) AngolanKwanza() AngolanKwanza {
	return AngolanKwanza(n * 1.08469)
}

func (n Naira) Euro() Euro {
	return Euro(n * 0.00229)
}

func (n Naira) BarbadosDollar() BarbadosDollar {
	return BarbadosDollar(n * 0.00456)
}
