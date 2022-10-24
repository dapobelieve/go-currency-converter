package tunedev

type Naira float32 // Nigerian Currency

func (n Naira) toCongoleseFranc() CongoleseFranc {
	return CongoleseFranc(n * 4.60525)
}

func (n Naira) toEthopianBirr() EthopianBirr {
	return EthopianBirr(n * 0.12039)
}

func (n Naira) toSouthAfricanRand() Rand {
	return Rand(n * 0.04141)
}

func (n Naira) toEgyptianPound() EgyptianPound {
	return EgyptianPound(n * 0.04481)
}

func (n Naira) toAlgerianDinar() AlgerianDinar {
	return AlgerianDinar(n * 0.32049)
}

func (n Naira) toWestAfricaFrancCFA() WestAfricaFranc {
	return WestAfricaFranc(n * 1.52113)
}

func (n Naira) toBotswanaPula() BotswanaPula {
	return BotswanaPula(n * 0.03042)
}

func (n Naira) toGuineaFranc() GuineaFranc {
	return GuineaFranc(n * 19.4674)
}

func (n Naira) toCapeVerdeEscudo() Escudo {
	return Escudo(n * 0.2557)
}

func (n Naira) toGhanaianNewCedi() Cedi {
	return Cedi(n * 0.02959)
}
