package hussain4real

type Naira float32

type QatariRial float32

type Bitcoin float32

type SudanesePound float32

type AngolanKwanza float32

type BosnianMark float32

type Euro float32

type SwedishKrona float32

type NorwegianKroner float32

type CanadianDollar float32

type BritishPound float32

// QatariRials conversion from naira to Qatari Rial
func QatariRials(n Naira) QatariRial {
	return QatariRial(n / 117.74)
}

// Bitcoins conversion from naira to Bitcoin
func Bitcoins(n Naira) Bitcoin {
	return Bitcoin(n / 8370940)
}

// SudanesePounds conversion from naira to Sudanese Pound
func SudanesePounds(n Naira) SudanesePound {
	return SudanesePound(n * 1.32834)
}

// AngolanKwanzas conversion from naira to Angolan Kwanza
func AngolanKwanzas(n Naira) AngolanKwanza {
	return AngolanKwanza(n * 0.98568)
}

// BosnianMarks conversion from naira to Bosnian Mark
func BosnianMarks(n Naira) BosnianMark {
	return BosnianMark(n * 0.00461)
}

// Euros conversion from naira to Euro
func Euros(n Naira) Euro {
	return Euro(n * 0.00236)
}

// Kronas conversion from naira to Swedish Króna
func Kronas(n Naira) SwedishKrona {
	return SwedishKrona(n * 0.02563)
}

// Kroners conversion from naira to Norwegian Kroner
func Kroners(n Naira) NorwegianKroner {
	return NorwegianKroner(n * 0.02512)
}

// CanadianDollars conversion from naira to Canadian Dollar
func CanadianDollars(n Naira) CanadianDollar {
	return CanadianDollar(n * 0.00319)
}

// BritishPounds conversion from naira to British Pound
func BritishPounds(n Naira) BritishPound {
	return BritishPound(n * 0.00207)
}
