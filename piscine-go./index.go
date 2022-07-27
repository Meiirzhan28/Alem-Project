package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	c := []rune(toFind)
	n := 0
	m := 0
	for range str {
		n++
	}
	for range c {
		m++
	}
	for i := 0; i <= n-m; i++ {
		if toFind == s[i:i+m] {
			return (i)
		}
	}
	return -1
}
