package pelumichael

type Naira float32

type Euro float32

type Pound float32

type CanadianDollar float32

type GhanianCedi float32

type IndianRupee float32

type KuwaitDinar float32

type KenyanShilling float32

type EgyptianPound float32

type CongoleseFranc float32

func EUR(n Naira) Euro {
	return Euro(n * 421.9457)
}

func GBP(n Naira) Pound {
	return Pound(n * 480.1814)
}

func CAD(n Naira) CanadianDollar {
	return CanadianDollar(n * 314.7864)
}

func GHS(n Naira) GhanianCedi {
	return GhanianCedi(n * 40.7701)
}

func INR(n Naira) IndianRupee {
	return IndianRupee(n * 5.2785)
}

func KWD(n Naira) KuwaitDinar {
	return KuwaitDinar(n * 1396.8206)
}

func KES(n Naira) KenyanShilling {
	return KenyanShilling(n * 3.5714)
}

func EGP(n Naira) EgyptianPound {
	return EgyptianPound(n * 22.0527)
}

func CDF(n Naira) CongoleseFranc {
	return CongoleseFranc(n * 0.2090)
}
