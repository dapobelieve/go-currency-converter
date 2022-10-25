package Emmrysjay

type NGN float32 // Nigerian Naira

type JPY float32 // Japanese Yen

type AFN float32 // Afghan Aghani

type ALL float32 // Albanian Lek

type AMD float32 // Armenian Dram

type BBD float32 // Barbados Dollar

type BYN float32 // Belarusian Ruble

type ISK float32 // Iceland Krona

type GNF float32 // Guinea Franc

type KRW float32 // Koream Won

type STN float32 // Sao Tome/Principe Dobra

func (n NGN) JPY() JPY { return JPY(n * 0.33423) }

func (n NGN) AFN() AFN { return AFN(n * 0.20009) }

func (n NGN) ALL() ALL { return ALL(n * 0.27108) }

func (n NGN) AMD() AMD { return AMD(n * 0.88273) }

func (n NGN) BBD() BBD { return BBD(n * 0.00462) }

func (n NGN) BYN() BYN { return BYN(n * 0.00585) }

func (n NGN) ISK() ISK { return ISK(n * 0.3273) }

func (n NGN) GNF() GNF { return GNF(n * 19.7233) }

func (n NGN) KRW() KRW { return KRW(n * 3.24614) }

func (n NGN) STN() STN { return STN(n * 0.05667) }
