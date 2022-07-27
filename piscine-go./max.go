package piscine

func Max(a []int) int {
	i := 1
	for i < len(a) {
		if a[i-1] > a[i] {
			tmp := a[i]
			a[i] = a[i-1]
			a[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
	return a[len(a)-1]
}
