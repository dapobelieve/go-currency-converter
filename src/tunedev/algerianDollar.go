package tunedev

type AlgerianDinar float32 // Algerian Dinar

func (ad AlgerianDinar) toNaira() Naira {
	return Naira(ad * 3.12027)
}
