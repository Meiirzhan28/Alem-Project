package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func average(s []int) float64 {
	a := 0
	for _, i := range s {
		a = a + i
	}
	return math.Round(float64(a / len(s)))
}

func Mediana(s []int) float64 {
	if len(s)%2 == 1 {
		a := (len(s) - 1) / 2
		return float64(s[a])
	} else {
		return float64((s[(len(s)/2)-1] + s[len(s)/2]) / 2)
	}
}

func variance(s []int) int {
	d := 0
	a := average(s)
	b := make([]float64, 0)
	for _, i := range s {
		b = append(b, (float64(i) - a))
	}
	for _, l := range b {
		d = d + int(math.Pow(float64(l), 2))
	}
	e := d / (len(b))
	return e
}

func deviation(s []int) float64 {
	d := 0
	a := average(s)
	b := make([]float64, 0)
	for _, i := range s {
		b = append(b, (float64(i) - a))
	}
	for _, l := range b {
		d = d + int(math.Pow(float64(l), 2))
	}
	e := d / (len(b))
	k := math.Sqrt(float64(e))
	return math.Round(k)
}

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	r := true
	j := make([]int, 0)
	input := string(file)
	text := strings.Fields(input)
	for j, i := range input {
		if i >= '0' && i <= '9' {
			if input[j+1] == 32 && input[j+2] >= '0' && input[j+2] <= '9' {
				r = false
			}
		}
	}
	if r == true {
		for i := range text {
			a, err := strconv.Atoi(text[i])
			if err != nil {
				log.Fatal(err)
			}
			j = append(j, a)

		}
		q := average(j)
		q1 := Mediana(j)
		q2 := variance(j)
		q3 := deviation(j)
		fmt.Println("Average:", q)
		fmt.Println("Median:", q1)
		fmt.Println("Variance:", q2)
		fmt.Println("Standard Deviation:", q3)
	} else {
		fmt.Println("ERRORRRRRRRRRRRRRRRRRRRRRRRRRRRR")
	}
}
