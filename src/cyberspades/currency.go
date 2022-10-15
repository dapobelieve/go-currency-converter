package cyberspades

type NGN float32 // Nigerian Naira
type EUR float32 // Euro
type RUB float32 // Russian Ruble
type GBP float32 // Great Britain Pound
type GHS float32 // Ghanian New Cedi
type JPY float32 // Japanese Yen
type AUD float32 // Australian Dollar
type AED float32 // United Arab Emirates Dirham
type CAD float32 // Canadian Dollar
type USD float32 // United States Dollar
type ARS float32 // Argentine Peso

// Naira to Euro
func (n NGN) EURConverter() EUR {
	return EUR(n * 0.00235)
}

// Naira to Russian Ruble
func (n NGN) RUBConverter() RUB {
	return RUB(n * 0.14489)
}

// Naira to Great Britain Pound
func (n NGN) GBPConverter() GBP {
	return GBP(n * 0.00204)
}

// Naira to Ghanian New Cedi
func (n NGN) GHSConverter() GHS {
	return GHS(n * 0.02487)
}

// Naira to Japanese Yen
func (n NGN) JPYConverter() JPY {
	return JPY(n * 0.33965)
}

// Naira to Australian Dollar
func (n NGN) AUDConverter() AUD {
	return AUD(n * 0.00366)
}

// Naira to United Arab Emirates Dirham
func (n NGN) AEDConverter() AED {
	return AED(n * 0.00843)
}

// Naira to Canadian Dollar
func (n NGN) CADConverter() CAD {
	return CAD(n * 0.00317)
}

// Naira to United States Dollar
func (n NGN) USDConverter() USD {
	return USD(n * 0.0023)
}

// Naira to Argentine Peso
func (n NGN) ARSConverter() ARS {
	return ARS(n * 0.34761)
}
