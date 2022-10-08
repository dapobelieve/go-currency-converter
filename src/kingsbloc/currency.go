package kingsbloc

type Naira float32 // Nigerian Naira

type AED float32 // Utd. Arab Emir. Dirham
type USD float32 // US Dollar
type MXN float32 // Mexican Peso
type GBP float32 // Great British Pound
type EUR float32 // Euro
type JPY float32 // Japanese Yen
type HKD float32 // Hong Kong Dollar
type INR float32 // Indian Rupee
type ZAR float32 // South African Rand
type NAD float32 // Namibia Dollar

func (n Naira) Dirham() AED {
	return AED(n * 0.00846)
}

func (n Naira) UsDollar() USD {
	return USD(n * 0.0023)
}

func (n Naira) MexicanPeso() MXN {
	return MXN(n * 0.04621)
}

func (n Naira) Pound() GBP {
	return GBP(n * 0.00207)
}

func (n Naira) Euro() EUR {
	return EUR(n * 0.00236)
}

func (n Naira) Yen() JPY {
	return JPY(n * 0.33429)
}

func (n Naira) HKDollar() HKD {
	return HKD(n * 0.01809)
}

func (n Naira) Rupee() INR {
	return INR(n * 0.18993)
}

func (n Naira) Rand() ZAR {
	return ZAR(n * 0.04156)
}

func (n Naira) NambianDollar() NAD {
	return NAD(n * 0.04156)
}
