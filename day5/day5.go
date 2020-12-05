package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	// Get program input
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	pattern := string(input)
	s := strings.Split(pattern, "\n")

	min := 9999
	max := 0
	total := 0
	totalSeats := 0
	for _, x := range s {
		i := strings.Replace(x, "F", "0", -1)
		i = strings.Replace(i, "B", "1", -1)
		i = strings.Replace(i, "L", "0", -1)
		i = strings.Replace(i, "R", "1", -1)
		v, _ := strconv.ParseInt(i, 2, 32)
		value := int(v)
		//row := value >> 3
		//col := value % 8
		//fmt.Println(row)
		//fmt.Println(col)

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		total += value
		totalSeats++
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Println("Total: ", total)
	fmt.Println("Total Seats: ", totalSeats)

	// Without iterating through all seats again, can we determine the missing value?

	// Sum all the numbers between the min and max values
	expected := ((max - min) + 1) * (min + max) / 2

	// Subtract the total of seat numbers, this should give us the missing number...our seat
	mySeat := expected - total
	fmt.Println("My seat: ", mySeat)

	//sort.Ints(allSeats)
	//fmt.Println("Sorted: ", allSeats)
	//var prev = 77
	//for _, v := range allSeats {
	//	if v - prev == 2 {
	//		fmt.Println("Your seat: ", v)
	//	}
	//	prev = v
	//}

}
