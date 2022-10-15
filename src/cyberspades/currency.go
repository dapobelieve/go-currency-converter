package cyberspades

type NGN float32	// Nigerian Naira
type EUR float32	// Euro
type RUB float32	// Russian Ruble
type GBP float32	// Great Britain Pound
type GHS float32	// Ghanian New Cedi
type JPY float32	// Japanese Yen
type AUD float32	// Australian Dollar
type AED float32	// United Arab Emirates Dirham
type CAD float32	// Canadian Dollar
type USD float32	// United States Dollar
type ARS float32	// Argentine Peso

// Naira to Euro
func (n Naira) EURConverter() EUR {
	return (n * 0.00235)
}

// Naira to Russian Ruble
func (n Naira) RUBConverter() RUB {
	return (n * 0.14489)
}

// Naira to Great Britain Pound
func (n Naira) GBPConverter GBP {
	return (n * 0.00204)
}

// Naira to Ghanian New Cedi
func (n Naira) GHSConverter GHS {
	return (n * 0.02487)
}

// Naira to Japanese Yen
func (n Naira) JPYConverter JPY {
	return (n * 0.33965)
}

// Naira to Australian Dollar
func (n Naira) AUDConverter AUD {
	return (n * 0.00366)
}

// Naira to United Arab Emirates Dirham
func (n Naira) AEDConverter AED {
	return (n * 0.00843)
}

// Naira to Canadian Dollar
func (n Naira) CADConverter CAD {
	return (n * 0.00317)
}
 
// Naira to United States Dollar
func (n Naira) USDConverter USD {
	return (n * 0.0023)
}

// Naira to Argentine Peso
func (n Naira) ARSConverter ARS {
	return (n * 0.34761)
}
