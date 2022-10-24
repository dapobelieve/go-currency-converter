package tunedev

type EgyptianPound float32 // Egyptian Pound

func (e EgyptianPound) toNaira() Naira {
	return Naira(e * 22.3171)
}
