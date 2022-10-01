package hussain4real

type Naira int

type Dollar int

func (n Naira) Dollars() Dollar {
	return Dollar(n * 650)
}
