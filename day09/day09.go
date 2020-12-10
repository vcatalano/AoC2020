package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

//const Preamble = 25
const SumToFind = 177777905
const MaxSeriesLen = 50

func isValid(preambleInts []int, number int) bool {
	for i := 0; i < len(preambleInts)-1; i++ {
		for j := i; j < len(preambleInts)-1; j++ {
			x := preambleInts[i]
			y := preambleInts[j+1]
			//fmt.Printf("%d %d\n", x, y)
			if x+y == number {
				return true
			}
		}
	}
	return false
}

func sum(series ...int) int {
	sum := 0
	for i := range series {
		sum += series[i]
	}
	return sum
}

func smallest(series ...int) int {
	result := math.MaxInt32
	for _, v := range series {
		if v < result {
			result = v
		}
	}
	return result
}

func largest(series ...int) int {
	result := math.MinInt32
	for _, v := range series {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(input), "\n")
	var numbers []int
	for _, line := range s {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	// Part 1
	//for i := 0; i < len(numbers) - Preamble; i++ {
	//	p := numbers[i : i+Preamble]
	//	n := numbers[i+Preamble]
	//	v := isValid(p, n)
	//	if !v {
	//		fmt.Printf("%d - %d - %t\n", p, n, v)
	//		return
	//	}
	//}

	for l := 2; l < MaxSeriesLen; l++ {
		fmt.Printf("Slice size: %d\n", l)
		for i := 0; i < len(numbers)-l; i++ {
			ss := numbers[i : i+l]
			if sum(ss...) == SumToFind {
				fmt.Printf("Found series: %v\n", ss)
				small := smallest(ss...)
				large := largest(ss...)
				fmt.Printf("Small: %d\n", small)
				fmt.Printf("Large: %d\n", large)
				fmt.Println(small + large)
			}
		}
	}
}
