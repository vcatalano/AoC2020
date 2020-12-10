package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
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

	re, _ := regexp.Compile(`(\d+)-(\d+)\s(.*):\s(.*)`)

	// Convert strings to integers
	var validCount = 0
	for _, line := range lines {
		result := re.FindStringSubmatch(line)

		min, _ := strconv.Atoi(result[1])
		max, _ := strconv.Atoi(result[2])
		character := result[3]
		password := result[4]

		f := string(password[min-1])
		s := string(password[max-1])

		if (f == character) != (s == character) {
			validCount++
		}
	}

	fmt.Println(validCount)
}
