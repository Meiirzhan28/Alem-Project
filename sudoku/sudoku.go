package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	problem := os.Args[1:]

	if Check(problem) == true {
		sqr := [9][9]rune{}
		sqr = construct(sqr, problem)

		if solution(&sqr) == true {
			for x := 0; x < 9; x++ {
				for y := 0; y < 9; y++ {
					if y != 8 {
						z01.PrintRune(rune(sqr[x][y]))
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(rune(sqr[x][y]))
					}
				}
				z01.PrintRune('\n')
			}
		} else {
			fmt.Println("Error")
		}
	}
}

// Check len sudoku
func Check(str []string) bool {
	for i := 0; i < len(str); i++ {
		if len(str[i]) != 9 {
			fmt.Println("Error")
			return false
		}
	}
	for i := 0; i < len(str); i++ {
		for _, d := range str[i] {
			if d == '/' || d == '0' {
				fmt.Println("Error")
				return false
			} else if d < 46 || d > 57 {
				fmt.Println("Error")
				return false
			}
		}
	}
	if len(str) != 9 {
		fmt.Println("Error")
		return false
	}
	return true
}

// create table
func construct(sqr [9][9]rune, str []string) [9][9]rune {
	for i := range str {
		for j := range str[i] {
			sqr[i][j] = rune(str[i][j])
		}
	}
	return sqr
}

// find dot
func Wheredot(sqr *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sqr[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// solutin
func solution(sqr *[9][9]rune) bool {
	if !Wheredot(sqr) {
		return true
	}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if sqr[x][y] == '.' {
				for b := '1'; b <= '9'; b++ {
					if valid(sqr, x, y, b) {
						sqr[x][y] = b
						if solution(sqr) {
							return true
						}
					}
					sqr[x][y] = '.'
				}
				return false
			}
		}
	}
	return false
}

// validation

func valid(sqr *[9][9]rune, x int, y int, u rune) bool {
	for i := 0; i < 9; i++ {
		if u == sqr[x][i] {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if u == sqr[j][y] {
			return false
		}
	}
	a := x / 3
	b := y / 3

	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if u == sqr[k][l] {
				return false
			}
		}
	}
	return true
}
