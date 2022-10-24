package tunedev

type Cedi float32 // Ghanaian New Cedi

func (c Cedi) toNaira() Naira {
	return Naira(c * 33.7487)
}
