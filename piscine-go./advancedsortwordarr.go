package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	SortWordArr(a)
	f(a[1], a[2])
}
