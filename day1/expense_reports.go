package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func find_2020_prod(values []int) int {
	for _, x := range values {
		for _, y := range values {
			if x+y == 2020 {
				return x * y
			}
		}
	}
	return 0
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	text := string(input)
	lines := strings.Split(text, "\n")

	// Remove empty lines
	temp := lines[:0]
	for _, x := range lines {
		if len(x) > 0 {
			temp = append(temp, x)
		}
	}
	lines = temp

	var ints []int
	for _, i := range lines {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ints = append(ints, j)
	}

	fmt.Println(find_2020_prod(ints))
}
