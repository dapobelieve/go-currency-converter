package seyifunmiSF

type NGN float32 //Nigeran naira
type JPY float32 //Japanese Yen
type USD float32 //United States Dollar
type AUD float32 //Australia Dollar
type GBP float32 //Great British Pound
type ARS float32 //Argentine peso
type BRL float32 //Brazillian Real
type CAD float32 //Canadian Dollar
type ZAR float32 //South African Rand
type EUR float32 //Euro
type AED float32 //United Arab Emirates

func (n NGN) JPY() JPY {
	return JPY(n * 0.34134)
}

func (n NGN) USD() USD {
	return USD(n * 0.0023)
}

func (n NGN) AUD() AUD {
	return AUD(n * 0.0037)
}

func (n NGN) GBP() GBP {
	return GBP(n * 0.00205)
}

func (n NGN) ARS() ARS {
	return ARS(n * 0.34808)
}

func (n NGN) BRL() BRL {
	return BRL(n * 0.01222)
}

func (n NGN) CAD() CAD {
	return CAD(n * 0.00319)
}

func (n NGN) ZAR() ZAR {
	return ZAR(n * 0.04208)
}

func (n NGN) EUR() EUR {
	return EUR(n * 0.00236)
}

func (n NGN) AED() AED {
	return AED(n * 0.00843)
}
