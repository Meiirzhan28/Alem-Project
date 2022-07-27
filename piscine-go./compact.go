package piscine

func Compact(ptr *[]string) int {
	a := *ptr
	count := 0
	for _, v := range *ptr {
		if v != "" {
			a[count] = v
			count++
		}
	}
	*ptr = a[0:count]
	return count
}
