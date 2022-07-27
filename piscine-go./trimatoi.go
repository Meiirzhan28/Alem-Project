package piscine

func TrimAtoi(s string) int {
	a := []rune(s)
	res := 0
	if s == "" {
		return 0
	}
	mode := 1
	for i := 0; i < len(a); i++ {
		if a[i] == '-' && res == 0 {
			mode = -1
		}
		if a[i] < '0' || a[i] > '9' {
			continue
		}
		res *= 10
		res += int(a[i] - '0')
	}

	return res * mode
}
