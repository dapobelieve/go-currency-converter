package ogbenioye

//CURRENCIES
// HKDollar ---> Hong Kong Dollar
// Yen ---> Japanese Yen
// Cedis ---> Ghanian Cedis
// Won ---> Korean Won
// Riyal ---> Qatari Riyal
// Rouble ---> Russian Rouble
// Shilling ---> Ugandan Shilling
// Rand ---> South African Rand
// JMD ---> Jamaican Dollar

type Naira float32
type Rupee float32
type HKDollar float32
type Yen float32
type Cedis float32
type Won float32
type Riyal float32
type Rouble float32
type Shilling float32
type Rand float32
type JMD float32

func (n Naira) Rupee() Rupee {
	return Rupee(n * 0.1886)
}

func (n Naira) HKDollar() HKDollar {
	return HKDollar(n * 0.01813)
}

func (n Naira) Yen() Yen {
	return Yen(n * 0.33444)
}

func (n Naira) Cedis() Cedis {
	return Cedis(n * 0.02397)
}

func (n Naira) Won() Won {
	return Won(n * 3.32163)
}

func (n Naira) Riyal() Riyal {
	return Riyal(n * 0.00829)
}

func (n Naira) Rouble() Rouble {
	return Rouble(n * 0.13485)
}

func (n Naira) Shilling() Shilling {
	return Shilling(n * 8.8115)
}

func (n Naira) Rand() Rand {
	return Rand(n * 0.04147)
}

func (n Naira) JMD() JMD {
	return JMD(n * 0.34624)
}
