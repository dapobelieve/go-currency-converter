package tunedev

type Rand float32 // SouthAfrican Rand

func (r Rand) toNaira() Naira {
	return Naira(r * 24.15)
}
