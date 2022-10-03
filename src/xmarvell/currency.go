package xmarvell

type NGR float32 // Nigerian Naira

type YEN float32 // Japanese Yen
type ETH float32 // Ethereum
type CDF float32 // Congolese Franc
type BRL float32 // Brazilian Real
type TRY float32 // Turkish Lira
type NOK float32 // Norwegian Kroner
type EGP float32 // Egyptian Pound
type MDL float32 // Moldovan Leu
type ILS float32 // Israeli New Shekel
type GBP float32 // Great British Pound

// YEN converts Nigerian Naira to Japanese Yen
func (n NGR) YEN() YEN {
	return YEN(n * 0.3343)
}

// ETH converts Nigerian Naira to Ethereum
func (n NGR) ETH() ETH {
	return ETH(n * 0.3343)
}

// CDF converts Nigerian Naira to Congolese Franc
func (n NGR) CDF() CDF {
	return CDF(n * 4.64112)
}

// BRL converts Nigerian Naira to Brazilian Real
func (n NGR) BRL() BRL {
	return BRL(n * 0.0125)
}

// TRY converts Nigerian Naira to Turkish Lira
func (n NGR) TRY() TRY {
	return TRY(n * 0.04258)
}

// NOK converts Nigerian Naira to Norwegian Kroner
func (n NGR) NOK() NOK {
	return NOK(n * 0.02511)
}

// EGP converts Nigerian Naira to Egyptian Pound
func (n NGR) EGP() EGP {
	return EGP(n * 0.04506)
}

// MDL converts Nigerian Naira to Moldovan Leu
func (n NGR) MDL() MDL {
	return MDL(n * 0.04458)
}

// ILS converts Nigerian Naira to Israeli New Shekel
func (n NGR) ILS() ILS {
	return ILS(n * 0.00821)
}

// GBP converts Nigerian Naira to Great British Pound
func (n NGR) GBP() GBP {
	return GBP(n * 0.00207)
}
