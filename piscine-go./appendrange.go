package piscine

func AppendRange(min, max int) []int {
	var a []int
	if max < min {
		return a
	}

	for i := min; i < max; i++ {
		a = append(a, i)
	}
	return a
}
