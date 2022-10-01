package dapobelieve

type Naira float32

type Dollar float32

func (n Naira) Dollars() Dollar {
	return Dollar(n * 650)
}
