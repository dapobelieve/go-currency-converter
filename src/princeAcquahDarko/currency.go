package princeacquahdarko

type Naira float32

type Cedis float32
type Dollars float32
type Euros float32
type Pounds float32
type Kwanza float32
type Guilder float32
type Dram float32
type Lek float32
type Afghani float32
type Dirham float32


func (n Naira) CediConverter() Cedis{
	return Cedis(n * 0.0238)
}

func (n Naira) DollarConverter() Dollars{
	return Dollars(n * 0.00231)
}

func (n Naira) EurosConverter() Euros{
	return Euros(n * 0.00236)
}

func (n Naira) PoundsConverter() Pounds{
	return Pounds(n * 0.00207)
}

func (n Naira) KwanzaConverter() Kwanza{
	return Kwanza(n * 0.98568)
}

func (n Naira) GuilderConverter() Guilder{
	return Guilder(n * 0.00414)
}

func (n Naira) DramConverter() Dram{
	return Dram(n * 0.90119)
}

func (n Naira) LekConverter() Lek{
	return Lek(n * 0.27367)
}

func (n Naira) AfghaniConverter() Afghani{
	return Afghani(n * 0.20242)
}

func (n Naira) DirhamConverter() Dirham{
	return Dirham(n * 0.00848)
}



