package piscine

func Map(f func(int) bool, a []int) []bool {
	counter := 0
	for range a {
		counter++
	}
	r := make([]bool, counter)
	for i, val := range a {
		r[i] = f(val)
	}
	return r
}
