package piscine

func Abort(a, b, c, d, e int) int {
	f := []int{a, b, c, d, e}
	ss(f)
	return f[2]
}

func ss(table []int) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
}
