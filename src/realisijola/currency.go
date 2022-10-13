package  realisijola


type Naira float32    // Nigerian Naira
type Ruble float32    // Russian Rubble
type Ringgit float32  // Malaysian Ringit
type Shilling float32 // Uganda 
type Franc float32    // Rwandan Franc
type Sol float32      // Peruvian Nuevo Sol
type Escudo float32   // Cape Verde Escudo
type Tdollar float32  // Trinidad/Tobago Dollar
type Srupee float32   // Seychelles Rupee
type Fdollar float32  // Fiji Dollar
type Ldollar float32  // Liberian Dollar


func (n Naira) Ruble() Ruble {
	return Ruble(n * 0.1359)
}
func (n Naira) Ringgit() Ringgit {
	return Ringgit(n * 0.01073)
}
func (n Naira) Shilling() Shilling {
	return Shilling(n * 8.8115)
}
func (n Naira) Franc() Franc {
	return Franc(n * 2.39397)
}

func (n Naira) Sol() Sol {
	return Sol(n * 0.00906)
}

func (n Naira) Tdollar() Tdollar {
	return Tdollar(n * 0.0153)
} 

func (n Naira) Srupee() Srupee {
	return Srupee(n * 0.02871)
} 

func (n Naira) Fdollar() Fdollar {
	return Fdollar(n * 0.00528)
} 

func (n Naira) Ldollar() Ldollar {
	return Ldollar(n * 0.35268)
} 
