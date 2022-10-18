package kelechiotu

type Naira float32

type Dollar float32

type Rupee float32

type BritishPound float32

type DominicanPeso float32

type BrazilianReal float32

type EthiopianBirr float32

type JapaneseYen float32

type GhanaianNewCedi float32

type KenyanShilling float32

type ArubanFlorin float32

func (n Naira) Dollars() Dollar {
	return Dollar(n * 434.676)
}

func (n Naira) Rupees() Rupee {
	return Rupee(n * 5.2744)
}

func (n Naira) BritishPounds() BritishPound {
	return BritishPound(n * 486.101)
}

func (n Naira) DominicanPesos() DominicanPeso {
	return DominicanPeso(n * 8.00706)
}

func (n Naira) BrazilianReals() BrazilianReal{
	return BrazilianReal(n * 81.5887)
}

func (n Naira) EthiopianBirrs() EthiopianBirr {
	return EthiopianBirr(n * 8.22387)
}

func (n Naira) JapaneseYens() JapaneseYen {
	return JapaneseYen(n * 2.9219)
}

func (n Naira) GhanaianNewCedis() GhanaianNewCedi {
	return GhanaianNewCedi(n * 39.1408)
}

func (n Naira) KenyanShillings() KenyanShilling {
	return KenyanShilling(n * 3.56293)
}

func (n Naira) ArubanFlorins() ArubanFlorin {
	return ArubanFlorin(n * 240.152)
}
