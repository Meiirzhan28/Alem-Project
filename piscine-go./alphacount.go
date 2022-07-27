package piscine

func AlphaCount(s string) int {
	a := 0
	for i := 0; i < len([]rune(s)); i++ {
		if []rune(s)[i] <= 'Z' && []rune(s)[i] >= 'A' || []rune(s)[i] <= 'z' && []rune(s)[i] >= 'a' {
			a++
		}
	}
	return a
}
