package piscine

func Unmatch(a []int) int {
	var quant int
	for _, el := range a {
		quant = 0
		for _, v := range a {
			if v == el {
				quant++
			}
		}
		if quant%2 != 0 {
			return el
		}
	}
	return -1
}
