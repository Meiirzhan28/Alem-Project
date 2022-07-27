package piscine

/*/func f(a, b int) int {
	if a > b {
		return 0
	} else if a == b {
		return a
	} else {
		return 1
	}
}
/*/

func IsSorted(f func(a, b int) int, a []int) bool {
	w := true
	g := true

	for k := 0; k < len(a)-1; k++ {
		if f(a[k], a[k+1]) > 0 {
			w = false
		}
	}

	for k := 0; k < len(a)-1; k++ {
		if f(a[k], a[k+1]) < 0 {
			g = false
		}
	}
	return g || w
}
