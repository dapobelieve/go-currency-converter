package tunedev

type Escudo float32 // Cape Verde Escudo

func (e Escudo) toNaira() Naira {
	return Naira(e * 3.91085)
}
