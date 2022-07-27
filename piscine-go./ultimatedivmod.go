package piscine

func UltimateDivMod(a *int, b *int) {
	z := *a
	*a = *a / *b
	*b = z % *b
}
