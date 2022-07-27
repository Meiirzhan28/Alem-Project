package piscine

func check(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	re := []rune(s)
	first := true
	for i := 0; i < len(re); i++ {
		if check(re[i]) && first {
			if re[i] >= 'a' && re[i] <= 'z' {
				re[i] -= 32
			}
			first = false
		} else if re[i] >= 'A' && re[i] <= 'Z' {
			re[i] += 32
		} else if !check(re[i]) {
			first = true
		}
	}
	return string(re)
}
