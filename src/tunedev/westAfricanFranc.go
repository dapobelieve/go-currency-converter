package tunedev

type WestAfricaFranc float32 // West Africa Franc

func (waf WestAfricaFranc) toNaira() Naira {
	return Naira(waf * 0.65741)
}
