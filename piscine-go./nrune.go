package piscine

func NRune(s string, n int) rune {
	res := []rune(s)
	for t, val := range res {
		if t == n-1 {
			return val
		}
	}
	return 0
}
