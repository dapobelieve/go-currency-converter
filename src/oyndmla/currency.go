package oyndmla

type Naira float32

type pound float32
type pesto float32
type dirham float32
type Guilder float32
type Florin float32
type euro float32
type lek float32
type pula float32
type afgani float32
type cedi float32


func (in naira) PoundConverter()Pound{
  return Pound(n * 0.00205)
}

func (in naira) PestoConverter()Pound{
  return pesto(n * 0.34405)
}

func (in naira) DirhamConverter()Pound{
  return Dirham(n * 0.00848)
}

func (in naira) GuilderConverter()Pound{
  return Guilder(n * 0.00413)
}

func (in naira) FlorinConverter()Pound{
  return Florin(n * 0.00205)
}

func (in naira) euroConverter()Pound{
  return euro(n * 0.00234)
}

func (in naira) lekConverter()Pound{
  return lek(n * 0.27108)
}

func (in naira) pulaConverter()Pound{
  return pula(n * 0.03016)
}

func (in naira) afghaniConverter()Pound{
  return Pound(n * 0.2009)
}

func (in naira) cediConverter()Pound{
  return cedi(n * 0.0241)
}
