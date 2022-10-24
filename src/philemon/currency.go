package philemon

type (
	// Naira is the Nigerian Naira
	Naira float64
	// BelizeDollar is the Belize Dollar.
	BelizeDollar float64
	// MoroccanDirham is the Moroccan Dirham.
	MoroccanDirham float64
	// GuineanFranc is the Guinean Franc.
	GuineanFranc float64
	// CanadianDollar is the Canadian Dollar.
	CanadianDollar float64
	// ChileanPeso is the Chilean Peso.
	ChileanPeso float64
	// LebanesePound is the Lebanese Pound.
	LebanesePound float64
	// ChineseYuan is the Chinese Yuan.
	ChineseYuan float64
	// UnitedStatesDollar is the United States Dollar.
	UnitedStatesDollar float64
	// ZimbabweanDollar is the Zimbabwean Dollar.
	ZimbabweanDollar float64
	// KoreanWon is the Korean Won.
	KoreanWon float64
)

// ToBelizeDollar converts a Naira value to Belize Dollar.
func (n Naira) ToBelizeDollar() BelizeDollar {
	return BelizeDollar(n / 218)
}

// ToMoroccanDirham converts a Naira value to Moroccan Dirham.
func (n Naira) ToMoroccanDirham() MoroccanDirham {
	return MoroccanDirham(n / 39.89)
}

// ToGuineanFranc converts a Naira value to Guinean Franc.
func (n Naira) ToGuineanFranc() GuineanFranc {
	return GuineanFranc(n / 19.72)
}

// ToCanadianDollar converts a Naira value to Canadian Dollar.
func (n Naira) ToCanadianDollar() CanadianDollar {
	return CanadianDollar(n / 319.79)
}

// ToChileanPeso converts a Naira value to Chilean Peso.
func (n Naira) ToChileanPeso() ChileanPeso {
	return ChileanPeso(n / 0.45)
}

// ToLebanesePound converts a Naira value to Lebanese Pound.
func (n Naira) ToLebanesePound() LebanesePound {
	return LebanesePound(n / 0.29)
}

// ToChineseYuan converts a Naira value to Chinese Yuan.
func (n Naira) ToChineseYuan() ChineseYuan {
	return ChineseYuan(n / 59.84)
}

// ToUnitedStatesDollar converts a Naira value to BelizeDollar.
func (n Naira) ToUnitedStatesDollar() UnitedStatesDollar {
	return UnitedStatesDollar(n / 437.18)
}

// ToZimbabweanDollar converts a Naira value to BelizeDollar.
func (n Naira) ToZimbabweanDollar() ZimbabweanDollar {
	return ZimbabweanDollar(n / 1.4)
}

// ToKoreanWon converts a Naira value to KoreanWon.
func (n Naira) ToKoreanWon() KoreanWon {
	return KoreanWon(n / 0.31)
}
