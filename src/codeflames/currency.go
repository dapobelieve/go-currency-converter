package codeflames

// Base currency from which all conversions are made.
type Naira float32 // Nigerian Naira

type EUR float32 // Euro
type JPY float32 // Japanese Yen
type BGN float32 // Bulgarian Lev
type GBP float32 // Pound Sterling
type PLN float32 // Polish Zloty
type SEK float32 // Swedish Krona
type GHS float32 // Ghanaian new cedi
type KES float32 // Kenyan Shilling
type ZAR float32 // South African Rand
type INR float32 // Indian Rupee

// Convert Naira to Euro
func (n Naira) Euro() EUR {
	return EUR(n * 0.0023)
}

// Convert Naira to Japanese Yen
func (n Naira) JapaneseYen() JPY {
	return JPY(n * 0.33792)
}

// Convert Naira to Bulgarian Lev
func (n Naira) BulgarianLev() BGN {
	return BGN(n * 0.00449)
}

// Convert Naira to Pound Sterling
func (n Naira) PoundSterling() GBP {
	return GBP(n * 0.00197)
}

// Convert Naira to Polish Zloty
func (n Naira) PolishZloty() PLN {
	return PLN(n * 0.01083)
}

// Convert Naira to Swedish Krona
func (n Naira) SwedishKrona() SEK {
	return SEK(n * 0.02505)
}

// Convert Naira to Ghanaian new Cedi
func (n Naira) GhanaianNewCedi() GHS {
	return GHS(n * 0.03156)
}

// Convert Naira to Kenyan Shilling
func (n Naira) KenyanShilling() KES {
	return KES(n * 0.27426)
}

// Convert Naira to South African Rand
func (n Naira) SouthAfricanRand() ZAR {
	return ZAR(n * 0.04164)
}

// Convert Naira to Indian Rupee
func (n Naira) IndianRupee() INR {
	return INR(n * 0.18814)
}
