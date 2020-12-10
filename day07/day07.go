package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	count int
	color string
}

// Given a rule string, parse it and create a rule object
func parseRule(line string) Rule {
	re := regexp.MustCompile(`^(\d+)\s(.+) bags?$`)
	// Remove the period at the end of the line and covert to fields
	fields := re.FindStringSubmatch(strings.Replace(line, ".", "", -1))

	// Special case for "no other bags"
	if len(fields) == 0 {
		return Rule{0, ""}
	}

	count, _ := strconv.Atoi(fields[1])
	return Rule{count, fields[2]}
}

func numberOfBags(allRules map[string][]Rule, bagRules []Rule) int {
	count := 0
	for _, br := range bagRules {
		count += br.count
		subBag := allRules[br.color]
		count += br.count * numberOfBags(allRules, subBag)
	}
	return count
}

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(input), "\n")
	fmt.Println(s)

	re := regexp.MustCompile(`^(.+)\s(bags contain)\s(.+)$`)
	allRules := make(map[string][]Rule)
	for _, line := range s {
		fields := re.FindStringSubmatch(line)

		rules := strings.Split(fields[3], ", ")
		var rule []Rule
		// Parse each of the rules for a given bag color and add them to the master list of rules
		for _, r := range rules {
			fmt.Println(parseRule(r))
			rule = append(rule, parseRule(r))
		}
		allRules[fields[1]] = rule
	}

	fmt.Println(numberOfBags(allRules, allRules["shiny gold"]))
}
