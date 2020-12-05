package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const PatternRepeats = 200

func countTrees(a [][]string, rows int, rightStep int, downStep int) int {
	treeCount := 0
	r := 0
	c := 0
	for r+downStep < rows {
		r += downStep
		c += rightStep
		if a[r][c] == "#" {
			treeCount++
		}
	}

	return treeCount
}

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	pattern := string(input)

	s := strings.Split(pattern, "\n")
	rows := len(s)

	a := make([][]string, rows)
	for i := range a {
		a[i] = strings.Split(strings.Repeat(s[i], PatternRepeats), "")
	}

	aa := countTrees(a, rows, 1, 1)
	bb := countTrees(a, rows, 3, 1)
	cc := countTrees(a, rows, 5, 1)
	dd := countTrees(a, rows, 7, 1)
	ee := countTrees(a, rows, 1, 2)

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	fmt.Println(aa * bb * cc * dd * ee)
}
