package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(input), "\n\n")

	re := regexp.MustCompile(`\s+`)
	sum := 0
	for _, groups := range s {
		fields := re.Split(groups, -1)

		fmt.Println(fields)
		peopleInGroup := len(fields)
		set := make(map[string]int)
		for _, f := range fields {
			a := strings.Split(f, "")
			for _, c := range a {
				fmt.Println(c)
				if count, ok := set[c]; ok {
					set[c] = count + 1
				} else {
					set[c] = 1
				}
			}
		}

		fmt.Println(set)

		// For each answer which all people in the group answered "yes", count this answer
		for _, v := range set {
			if v == peopleInGroup {
				sum++
			}
		}
		fmt.Println()
	}
	fmt.Println(sum)
}
