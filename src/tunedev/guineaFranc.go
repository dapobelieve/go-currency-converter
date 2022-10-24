package tunedev

type GuineaFranc float32 // Guinea Franc

func (gf GuineaFranc) toNaira() Naira {
	return Naira(gf * 0.05137)
}
