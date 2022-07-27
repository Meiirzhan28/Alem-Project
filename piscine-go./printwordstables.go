package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		strRune := []rune(v)
		for _, d := range strRune {
			z01.PrintRune(d)
		}
		z01.PrintRune('\n')
	}
}
