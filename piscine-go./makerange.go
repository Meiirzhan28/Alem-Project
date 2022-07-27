package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return []int(nil)
	}
	s := max - min
	a := make([]int, s)
	for i := 0; i < s; i++ {
		a[i] = i + min
	}
	return a
}
