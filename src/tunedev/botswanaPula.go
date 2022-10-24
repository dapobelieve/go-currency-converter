package tunedev

type BotswanaPula float32 // Botswana Pula

func (b BotswanaPula) toNaira() Naira {
	return Naira(b * 32.8747)
}
