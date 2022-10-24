package tunedev

type CongoleseFranc float32 // Congolese Franc

func (c CongoleseFranc) toNaira() Naira {
	return Naira(c * 0.21714)
}
