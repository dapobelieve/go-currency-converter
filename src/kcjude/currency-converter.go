package kcjude

import "fmt"

type nigeriaNaira float32

type belizeDollar float32

type indonesiaRupiah float32

type mauritiusRupee float32

type indianRupee float32

type lebanesePound float32

type lesothoLoti float32

type algeriaDinar float32

type czechKoruna float32

type qatariRia float32

type zambiaKwacha float32

func main() {
zambia := nigeriaNaira.zambiaKwacha(10)
fmt.Println(zambia)
}

func (n nigeriaNaira) belizeDollar() belizeDollar {
	return belizeDollar(n * 0.0046)
}

func (n nigeriaNaira) indonesiaRupiah() indonesiaRupiah {
	return indonesiaRupiah(n * 35.52)
}

func (n nigeriaNaira) mauritiusRupee() mauritiusRupee {
	return mauritiusRupee(n * 0.10)
}

func (n nigeriaNaira) indianRupee() indianRupee {
	return indianRupee(n * 0.19)
}

func (n nigeriaNaira) lebanesePound() lebanesePound {
	return lebanesePound(n * 3.47)
}

func (n nigeriaNaira) lesothoLoti() lesothoLoti {
	return lesothoLoti(n * 0.041)
}

func (n nigeriaNaira) algeriaDinar() algeriaDinar {
	return algeriaDinar(n * 0.32)
}

func (n nigeriaNaira) czechKoruna() czechKoruna {
	return czechKoruna(n * 0.056)
}

func (n nigeriaNaira) qatariRial() qatariRia {
	return qatariRia(n * 0.0083)
}

func (n nigeriaNaira) zambiaKwacha() zambiaKwacha {
	return zambiaKwacha(n * 0.037)
}
