package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func isValid(adapters []int) bool {
	prev := 0
	for _, a := range adapters {
		if a-prev > 3 {
			return false
		}
		prev = a
	}
	return true
}

func sum(series ...int) int {
	sum := 0
	for i := range series {
		sum += series[i]
	}
	return sum
}

// I don't have time to figure out memoization with Golang at this time, instead, I have a
// lookup of adapter sub-slices based on the sum of the sub slices (the sum is probably
// guaranteed to be unique for each sub-slice)
var countLookup = make(map[int]int)

func countForSubset(first int, adapters []int) int {

	// Base rule
	if len(adapters) == 1 {
		return 1
	}

	// Look up the existing value for the count of a subset, this prevents us from
	// having to re-calculate whole sub-slices again (essentially, memoization)
	s := sum(adapters...)
	count := 0
	if value, exist := countLookup[s]; exist {
		count = value
	} else {
		count = countForSubset(adapters[0], adapters[1:])
		countLookup[s] = count
	}

	// Remove the first item and attempt to get the count for the rest
	if len(adapters) > 2 && (adapters[1]-first <= 3) {
		count += countForSubset(first, adapters[1:])
	}
	//fmt.Printf("Count: %d\n", count)
	return count
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
	var adapters []int
	for _, line := range s {
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)
	l := largest(adapters...) + 3
	adapters = append(adapters, l)
	fmt.Println(adapters)

	// Part 1
	//bucket1 := 0
	//bucket2 := 0
	//bucket3 := 0
	//prev := 0
	//
	//for _, a := range adapters {
	//	if a - prev == 1 {
	//		bucket1++
	//	} else if a - prev == 2 {
	//		bucket2++
	//	} else if a - prev == 3 {
	//		bucket3++
	//	} else {
	//		break
	//	}
	//	prev = a
	//}
	//
	//fmt.Printf("bucket1: %d\n", bucket1)
	//fmt.Printf("bucket2: %d\n", bucket2)
	//fmt.Printf("bucket3: %d\n", bucket3)
	//fmt.Println(bucket1 * bucket3)

	// Part 2
	start := time.Now()
	count := countForSubset(0, adapters)
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(count)

}
