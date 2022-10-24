package tunedev

type EthopianBirr float32 // Ethopian Birr

func (e EthopianBirr) toNaira() Naira {
	return Naira(e * 8.30613)
}
