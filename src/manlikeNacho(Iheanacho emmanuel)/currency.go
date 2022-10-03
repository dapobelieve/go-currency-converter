// Package manlikeNacho_Iheanacho_emmanuel_ converts Nigerian currency to various other currencies.
package manlikeNacho_Iheanacho_emmanuel_

type Naira float32
type Cedi float32
type CanadianDollar float32
type Real float32
type Yuan float32
type Kuna float32
type Peso float32
type Koruna float32
type Yen float32
type Dirham float32

// Cedi converts naira to ghanaian cedi.
func (n Naira) Cedi() Cedi {
	return Cedi(n / 41.2)
}

// Naira converts cedi to naira.
func (c Cedi) Naira() Naira {
	return Naira(c * 41.2)
}

// CanadianDollar converts naira to canadian dollars.
func (n Naira) CanadianDollar() CanadianDollar {
	return CanadianDollar(n / 314.6)
}

// Naira converts canadian dollars to naira.
func (c CanadianDollar) Naira() Naira {
	return Naira(c * 314.6)
}

// Real converts naira to brazilian real.
func (n Naira) Real() Real {
	return Real(n * 0.013)
}

// Naira converts real to naira.
func (r Real) Naira() Naira {
	return Naira(r / 79.8)
}

// Yuan converts naira to yuan.
func (n Naira) Yuan() Yuan {
	return Yuan(n * 0.016)
}

// Naira converts Yuan to Naira.
func (y Yuan) Naira() Naira {
	return Naira(y / 0.016)
}

// Kuna converts naira to Crotian Kuna.
func (n Naira) Kuna() Kuna {
	return Kuna(n * 0.018)
}

// Naira converts Kuna to naira.
func (k Kuna) Naira() Naira {
	return Naira(k / 0.018)
}

// Peso converts naira to Cuban Peso.
func (n Naira) Peso() Peso {
	return Peso(n * 0.055)
}

// Naira converts cuban peso to Naira.
func (p Peso) Naira() Naira {
	return Naira(p / 0.055)
}

// Koruna converts naira to Czech Koruna.
func (n Naira) Koruna() Koruna {
	return Koruna(n * 0.058)
}

// Naira converts Czech Koruna to naira.
func (k Koruna) Naira() Naira {
	return Naira(k / 0.058)
}

// Yen converts naira to Japanese Yen.
func (n Naira) Yen() Yen {
	return Yen(n * 0.35)
}

// Naira converts Japanese yen to naira.
func (y Yen) Naira() Naira {
	return Naira(y / 0.35)
}

// Dirham converts naira to UAE dirham.
func (n Naira) Dirham() Dirham {
	return Dirham(n * 0.0085)
}

func (d Dirham) Naira() Naira {
	return Naira(d / 0.0085)
}
