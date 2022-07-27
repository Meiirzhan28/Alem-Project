package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func firstrune(s string) bool {
	if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0]=='h' {
		return true
	} else if s[0] == 'A' || s[0] == 'E' || s[0] == 'I' || s[0] == 'O' || s[0] == 'U' || s[0]=='H' {
		return true
	}
	return false
}

func lastrune(s string) bool {
	if s[len(s)-1] == ',' {
		return true
	}
	return false
}

func check(s string) string {
	if lastrune(s) == true {
		return s
	} else {
		a := strings.SplitAfter(s, ",")
		return strings.Join(a, " ")
	}
}

func proverka(s string) string {
	a := ""
	for j, i := range s {
		if i == ',' {
			a = a + s[:j+1]
		}
	}
	return a
}

func finddoremove(s []string) []string {
	find := regexp.MustCompile(`\w?\(up\)[[:punct:]]?|\w?\(cap\)[[:punct:]]?|\w?\(low\)[[:punct:]]?`)
	find2 := regexp.MustCompile(`\w?\(up\,\s*\d?[[:graph:]]?|\w?\(cap\,\s*\d?[[:graph:]]?|\w?\(low\,\s*\d?[[:graph:]]?`)
	a := 1
	for j := range s {
		if s[j] == "(hex)" {
			decimal, err := strconv.ParseInt(s[j-1], 16, 32)
			if err != nil {
				fmt.Println(err)
			}
			s[j-1] = fmt.Sprint(decimal)
			s[j] = ""
		}
		if s[j] == "(bin)" {
			decimal, err := strconv.ParseInt(s[j-1], 2, 64)
			if err != nil {
				fmt.Println(err)
			}
			s[j-1] = fmt.Sprint(decimal)
			s[j] = ""
		}
		if find.MatchString(s[j]) == true && s[j] == "(up)" || find2.MatchString(s[j]) == true && proverka(s[j]) == "(up," {
			if a == 1 {
				if j == 0 && s[j] == "(up)" && len(s) == 1 {
					s[j] = ""
					fmt.Println("Please Enter word before (up)")
				} else if j == 0 && s[j] == "(up)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(up," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (up)")
				} else if j == 0 && s[j] == "(up," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(up)" {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToUpper(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(up," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(up," {
							s[j] = check(s[j])
							s[j-1] = strings.ToUpper(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToUpper(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToUpper(s[j-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToUpper(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(up," {
							s[j] = check(s[j])
							s[j-1] = strings.ToUpper(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToUpper(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToUpper(s[j-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			} else if a > 1 {
				if j == 0 && s[j] == "(up)" && len(s) == 1 {
					fmt.Println("Please Enter word before (up)")
				} else if j == 0 && s[j] == "(up)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(up," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (up)")
				} else if j == 0 && s[j] == "(up," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(up)" {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToUpper(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(up," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(up," {
							s[j] = check(s[j])
							s[j-a] = strings.ToUpper(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-a-i] = strings.ToUpper(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-i-a] = strings.ToUpper(s[j-i-a])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToUpper(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(up," {
							s[j] = check(s[j])
							s[j-a] = strings.ToUpper(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-a-i] = strings.ToUpper(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-i-a] = strings.ToUpper(s[j-i-a])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			}
		} else if find.MatchString(s[j]) == true && s[j] == "(cap)" || find2.MatchString(s[j]) == true && proverka(s[j]) == "(cap," {
			if a == 1 {
				if j == 0 && s[j] == "(cap)" && len(s) == 1 {
					fmt.Println("Please Enter word before (cap)")
				} else if j == 0 && s[j] == "(cap)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(cap," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (cap)")
				} else if j == 0 && s[j] == "(cap," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(cap)" {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToLower(s[j-1])
							s[j-1] = strings.Title(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(cap," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(cap," {
							s[j] = check(s[j])
							s[j-1] = strings.ToLower(s[j-1])
							s[j-1] = strings.Title(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j-i] = strings.Title(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j-i] = strings.Title(s[j-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToLower(s[j-1])
							s[j-1] = strings.Title(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(cap," {
							s[j] = check(s[j])
							s[j-1] = strings.ToLower(s[j-1])
							s[j-1] = strings.Title(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j-i] = strings.Title(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j-i] = strings.Title(s[j-i])
									s[j+1] = ""
								}
							}
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			} else if a > 1 {
				if j == 0 && s[j] == "(cap)" && len(s) == 1 {
					fmt.Println("Please Enter word before (cap)")
				} else if j == 0 && s[j] == "(cap)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(cap," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (cap)")
				} else if j == 0 && s[j] == "(cap," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(cap)" {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToLower(s[j-a])
							s[j-a] = strings.Title(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(cap," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(cap," {
							s[j] = check(s[j])
							s[j-a] = strings.ToLower(s[j-a])
							s[j-a] = strings.Title(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-i-a] = strings.ToLower(s[j-i-a])
									s[j-a-i] = strings.Title(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-i-a] = strings.ToLower(s[j-i-a])
									s[j-i-a] = strings.Title(s[j-i-a])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToLower(s[j-a])
							s[j-a] = strings.Title(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(cap," {
							s[j] = check(s[j])
							s[j-a] = strings.ToLower(s[j-a])
							s[j-a] = strings.Title(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-a-i] = strings.ToLower(s[j-a-i])
									s[j-a-i] = strings.Title(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-a-i] = strings.ToLower(s[j-a-i])
									s[j-a-i] = strings.Title(s[j-a-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			}
		} else if find.MatchString(s[j]) == true && s[j] == "(low)" || find2.MatchString(s[j]) == true && proverka(s[j]) == "(low," {
			if a == 1 {
				if j == 0 && s[j] == "(low)" && len(s) == 1 {
					fmt.Println("Please Enter word before (low)")
				} else if j == 0 && s[j] == "(low)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(low," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (low)")
				} else if j == 0 && s[j] == "(low," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(low)" {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToLower(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(low," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(low," {
							s[j] = check(s[j])
							s[j-1] = strings.ToLower(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-1] = strings.ToLower(s[j-1])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(low," {
							s[j] = check(s[j])
							s[j-1] = strings.ToLower(s[j-1])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i <= b; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j; i++ {
									s[j-i] = strings.ToLower(s[j-i])
									s[j+1] = ""
								}
							}
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			} else if a > 1 {
				if j == 0 && s[j] == "(low)" && len(s) == 1 {
					fmt.Println("Please Enter word before (low)")
				} else if j == 0 && s[j] == "(low)" && len(s) > 1 {
					s[j] = ""
				} else if j == 0 && s[j] == "(low," && len(s) == 2 {
					s[j] = ""
					s[j+1] = ""
					fmt.Println("Please Enter word before (low)")
				} else if j == 0 && s[j] == "(low," && len(s) > 2 {
					s[j] = ""
					s[j+1] = ""
				} else {
					if j == len(s)-1 && s[j] == "(low)" {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToLower(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
						}
					} else if j == len(s)-2 && s[j] == "(low," {
						if find2.MatchString(s[j]) == true && proverka(s[j]) == "(low," {
							s[j] = check(s[j])
							s[j-a] = strings.ToLower(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-a-i] = strings.ToLower(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-i-a] = strings.ToLower(s[j-i-a])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
						}
					} else {
						if find.MatchString(s[j]) == true {
							s[j-a] = strings.ToLower(s[j-a])
							s[j] = find.ReplaceAllString(s[j], "")
							if find.MatchString(s[j+1]) == true || find2.MatchString(s[j+1]) == true {
								a += 1
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						} else if find2.MatchString(s[j]) == true && proverka(s[j]) == "(low," {
							s[j] = check(s[j])
							s[j-a] = strings.Title(s[j-a])
							s[j] = find2.ReplaceAllString(s[j], "")
							l := len(s[j+1])
							numb := s[j+1][:l-1]
							b, err := strconv.Atoi(numb)
							if err != nil {
								fmt.Println(err)
							}
							if b < j {
								for i := 0; i < b; i++ {
									s[j-a-i] = strings.ToLower(s[j-a-i])
									s[j+1] = ""
								}
							} else if b >= j {
								for i := 0; i <= j-a; i++ {
									s[j-a-i] = strings.ToLower(s[j-a-i])
									s[j+1] = ""
								}
							}
							s[j] = find2.ReplaceAllString(s[j], "")
							if find2.MatchString(s[j+2]) == true || find.MatchString(s[j+2]) == true {
								a += 2
								s[j] = find.ReplaceAllString(s[j], "")
							} else {
								a = 1
							}
						}
					}
				}
			}
		}
	}
	return s
}

func punt(input string) string {
	quotes := regexp.MustCompile(`\s{0,}\'\s{1,}`)
	quotes2 := regexp.MustCompile(`\s{1,}\'\s{0,}`)
	input = quotes.ReplaceAllString(input, " ' ")
	input = quotes2.ReplaceAllString(input, " ' ")
	morepunct := regexp.MustCompile(`\.\.\.`)
	morepunct2 := regexp.MustCompile(`\!\?`)
	morepunct3 := regexp.MustCompile(`\s*?\.\.\.\s*?`)
	morepunct4 := regexp.MustCompile(`\s*?\!\?\s*?`)
	input = morepunct.ReplaceAllString(input, "... ")
	input = morepunct2.ReplaceAllString(input, "!? ")
	input = morepunct3.ReplaceAllString(input, "...")
	input = morepunct4.ReplaceAllString(input, "!?")
	space := regexp.MustCompile(`\,`)
	space2 := regexp.MustCompile(`\.`)
	space3 := regexp.MustCompile(`\!`)
	space4 := regexp.MustCompile(`\?`)
	space5 := regexp.MustCompile(`\:`)
	space6 := regexp.MustCompile(`\;`)
	input = space.ReplaceAllString(input, ", ")
	input = space2.ReplaceAllString(input, ". ")
	input = space3.ReplaceAllString(input, "! ")
	input = space4.ReplaceAllString(input, "? ")
	input = space5.ReplaceAllString(input, ": ")
	input = space6.ReplaceAllString(input, "; ")
	punct := regexp.MustCompile(`\s*?\,\s*?`)
	punct2 := regexp.MustCompile(`\s*?\.\s*?`)
	punct3 := regexp.MustCompile(`\s*?\?\s*?`)
	punct4 := regexp.MustCompile(`\s*?\!\s*?`)
	punct5 := regexp.MustCompile(`\s*?\:\s*?`)
	punct6 := regexp.MustCompile(`\s*?\;\s*?`)
	input = punct.ReplaceAllString(input, ",")
	input = punct2.ReplaceAllString(input, ".")
	input = punct3.ReplaceAllString(input, "?")
	input = punct4.ReplaceAllString(input, "!")
	input = punct5.ReplaceAllString(input, ":")
	input = punct6.ReplaceAllString(input, ";")
	bracket := regexp.MustCompile(`\s*?\)`)
	bracket2 := regexp.MustCompile(`\)\,`)
	bracket3 := regexp.MustCompile(`\)\.`)
	bracket4 := regexp.MustCompile(`\)\!`)
	bracket5 := regexp.MustCompile(`\)\?`)
	bracket6 := regexp.MustCompile(`\)\:`)
	bracket7 := regexp.MustCompile(`\)\;`)
	bracket8 := regexp.MustCompile(`\)`)
	bracket9 := regexp.MustCompile(`\(`)
	input = bracket.ReplaceAllString(input, ")")
	input = bracket2.ReplaceAllString(input, ") ,")
	input = bracket3.ReplaceAllString(input, ") .")
	input = bracket4.ReplaceAllString(input, ") !")
	input = bracket5.ReplaceAllString(input, ") ?")
	input = bracket6.ReplaceAllString(input, ") :")
	input = bracket7.ReplaceAllString(input, ") ;")
	input = bracket8.ReplaceAllString(input, ") ")
	input = bracket9.ReplaceAllString(input, " (")
	doublespace := regexp.MustCompile(`\s{2,}`)
	input = doublespace.ReplaceAllString(input, " ")
	return input
}

func vowels(s []string) []string {
	vowels := regexp.MustCompile(`^[a]$`)
	vowels2 := regexp.MustCompile(`^[A]$`)
	for i := range s {
		if i == len(s)-1 {
			return s
		} else {
			if vowels.MatchString(s[i]) == true && firstrune(s[i+1]) == true {
				s[i] = vowels.ReplaceAllString(s[i], "an")
			} else if vowels2.MatchString(s[i]) == true && firstrune(s[i+1]) == true {
				s[i] = vowels2.ReplaceAllString(s[i], "An")
			}
		}
	}
	return s
}

func quotes2(s string) string {
	str := ""
	var removeSpace bool
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}

func quotes(s string) string {
	re := regexp.MustCompile(`\s{1,}\'\s{1,}(.*?)\s{1,}\'\s{1,}|\w??\'\s{1,}(.*?)\s{1,}\'\s{1,}`)
	b := re.FindAllString(s, -1)
	c := make([]string, 0)
	for i := 0; i < len(b); i++ {
		c = append(c, b[i])
	}
	g := strings.Fields(strings.Join(c, " "))
	for i := range g {
		if g[0] == "'" {
			g[1] = g[0] + g[1]
			g[0] = ""
		} else if g[i] == "'" {
			g[len(g)-2] = g[len(g)-2] + g[len(g)-1]
			g[len(g)-1] = ""
		}
	}
	s = strings.Join(g, " ")
	return s
}

func main() {
	if len(os.Args) == 3 {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		re := regexp.MustCompile(`\s{1,}\'\s{1,}(.*?)\s{1,}\'\s{1,}|\w??\'\s{1,}(.*?)\s{1,}\'\s{1,}`)
		input1 := string(file)
		q1 := strings.Fields(input1)
		result1 := make([]string, 0)
		for _, j := range q1 {
			result1 = append(result1, j)
		}
		q1 = vowels(result1)
		q2 := strings.Join(q1, " ")
		q3 := strings.Join(strings.Fields(q2), " ")

		input := punt(q3)
		q := quotes(input)
		text := strings.Fields(input)
		result := make([]string, 0)
		for _, j := range text {
			result = append(result, j)
		}
		result = finddoremove(result)
		result = vowels(result)
		a := strings.Join(result, " ")
		b := strings.Join(strings.Fields(a), " ")
		b = punt(b)
		b = re.ReplaceAllString(b, q)
		output := []byte(b)
		OurData := os.Args[2]
		os.WriteFile(OurData, output, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
