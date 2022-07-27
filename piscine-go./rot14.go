package piscine

func Rot14(s string) string {
	st := []rune(s)
	var str string
	for i := 0; i < len(st); i++ {
		if st[i] >= 'a' && st[i] <= 'l' {
			st[i] = st[i] + 14
		} else if st[i] >= 'A' && st[i] <= 'L' {
			st[i] = st[i] + 14
		} else if st[i] >= 'm' && st[i] <= 'z' {
			st[i] = st[i] - 12
		} else if st[i] >= 'M' && st[i] <= 'Z' {
			st[i] = st[i] - 12
		}
		str += string(st[i])
	}
	return str
}
