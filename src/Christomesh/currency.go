package Christomesh

type NGN float64

type BTC float64 // Bitcoin
type ETH float64 // Ethereum
type GBP float64 // Great British Pound
type YEN float64 // Japanese Yen
type ARS float64 // Argentine Peso
type CAD float64 // Canandian Dollar
type AED float64 // Utd Arab Emir Dirham
type AFN float64 // Afghna Afghani
type BDT float64 // Bangladeshi Taka
type AUD float64 // Australian Dollar

func (n NGN) BTC() BTC {
	return BTC(n / 8397293)
}

func (n NGN) ETH() ETH {
	return ETH(n / 574277)
}

func (n NGN) GBP() GBP {
	return GBP(n / 483.61)
}

func (n NGN) YEN() YEN {
	return YEN(n / 2.99132)
}

func (n NGN) ARS() ARS {
	return ARS(n / 2.93889)
}

func (n NGN) CAD() CAD {
	return CAD(n / 313.205)
}

func (n NGN) AED() AED {
	return AED(n / 117.872)
}

func (n NGN) AFN() AFN {
	return AFN(n / 4.94015)
}

func (n NGN) BDT() BDT {
	return BDT(n / 4.35227)
}

func (n NGN) AUD() AUD {
	return AUD(n / 277.607)
}
