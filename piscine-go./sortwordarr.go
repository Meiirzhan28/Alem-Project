package piscine

func SortWordArr(a []string) {
	for i := len(a) - 1; i > 0; i-- {
		for j := len(a) - 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
