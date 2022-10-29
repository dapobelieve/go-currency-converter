package clintonMF

type Naira float32 //Nigerian naira

type AED float32 // United arab emirates dirham
type CHF float32 //Swiss Franc
type GBP float32 //Great British pound
type JPY float32 //Japanese Yen
type PHP float32 //Philippine Peso
type SGD float32 //Singapore Dollar
type ZAR float32 //South African Rand
type AUD float32 //Australian Dollar
type CNY float32 //Chinese Yuan
type HKD float32 //Hong kong Dollar

func (n Naira) AED() AED {
	return AED(n * 0.00848)
}

func (n Naira) CHF() CHF {
	return CHF(n * 0.00227)
}

func (n Naira) GBP() GBP {
	return GBP(n * 0.00205)
}

func (n Naira) JPY() JPY {
	return JPY(n * 0.00848)
}

func (n Naira) PHP() PHP {
	return PHP(n * 0.1356)
}

func (n Naira) SGD() SGD {
	return SGD(n * 0.00329)
}

func (n Naira) ZAR() ZAR {
	return ZAR(n * 0.0412)
}

func (n Naira) AUD() AUD {
	return AUD(n * 0.00357)
}

func (n Naira) CNY() CNY {
	return CNY(n * 0.01642)
}

func (n Naira) HKD() HKD {
	return HKD(n * 0.01812)
}
