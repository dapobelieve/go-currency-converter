package dapobelieve

type congoleseFranc float32

type SouthAfricanRand float32

func (r SouthAfricanRand) Cdf() congoleseFranc {
	return congoleseFranc(r * 110.128)
}

func (f congoleseFranc) Zar() SouthAfricanRand {
	return SouthAfricanRand(f * 0.00876)
}
