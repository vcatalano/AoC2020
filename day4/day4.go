package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidBirthYear(val string) bool {
	matched, _ := regexp.MatchString(`^\d{4}$`, val)
	if !matched {
		return false
	}
	year, _ := strconv.Atoi(val)
	return 1920 <= year && year <= 2002
}

func isValidIssYear(val string) bool {
	matched, _ := regexp.MatchString(`^\d{4}$`, val)
	if !matched {
		return false
	}
	year, _ := strconv.Atoi(val)
	return 2010 <= year && year <= 2020
}

func isValidExpYear(val string) bool {
	matched, _ := regexp.MatchString(`^\d{4}$`, val)
	if !matched {
		return false
	}
	year, _ := strconv.Atoi(val)
	return 2020 <= year && year <= 2030
}

func isValidHeight(val string) bool {
	re, _ := regexp.Compile(`^(\d+)(cm|in)$`)
	result := re.FindStringSubmatch(val)

	if len(result) != 3 {
		return false
	}

	if result[2] == "in" {
		height, _ := strconv.Atoi(result[1])
		return 59 <= height && height <= 76
	} else if result[2] == "cm" {
		height, _ := strconv.Atoi(result[1])
		return 150 <= height && height <= 193
	}

	return false
}

func isValidHairColor(val string) bool {
	matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, val)
	return matched
}

func isValidEyeColor(val string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range colors {
		if color == val {
			return true
		}
	}
	return false
}

func isValidPid(val string) bool {
	matched, _ := regexp.MatchString(`^[0-9]{9}$`, val)
	return matched
}

func isValidPassport(passport map[string]string) bool {
	byr, ok := passport["byr"]
	if !ok || !isValidBirthYear(byr) {
		fmt.Printf("Invalid byr `%s`\n", byr)
		return false
	}

	iyr, ok := passport["iyr"]
	if !ok || !isValidIssYear(iyr) {
		fmt.Printf("Invalid iyr `%s`\n", iyr)
		return false
	}

	eyr, ok := passport["eyr"]
	if !ok || !isValidExpYear(eyr) {
		fmt.Printf("Invalid eyr `%s`\n", eyr)
		return false
	}

	hgt, ok := passport["hgt"]
	if !ok || !isValidHeight(hgt) {
		fmt.Printf("Invalid hgt `%s`\n", hgt)
		return false
	}

	hcl, ok := passport["hcl"]
	if !ok || !isValidHairColor(hcl) {
		fmt.Printf("Invalid hcl `%s`\n", hcl)
		return false
	}

	ecl, ok := passport["ecl"]
	if !ok || !isValidEyeColor(ecl) {
		fmt.Printf("Invalid ecl `%s`\n", ecl)
		return false
	}

	pid, ok := passport["pid"]
	if !ok || !isValidPid(pid) {
		fmt.Printf("Invalid pid `%s`\n", pid)
		return false
	}

	return true
}

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(input), "\n\n")

	re := regexp.MustCompile(`\s+`)
	validPassportCount := 0
	for _, x := range s {
		passportMap := make(map[string]string)
		fields := re.Split(x, -1)
		for _, f := range fields {
			v := strings.Split(f, ":")
			passportMap[v[0]] = v[1]
		}
		//fmt.Println(passportMap)
		if isValidPassport(passportMap) {
			fmt.Println(passportMap)
			validPassportCount++
		}
	}
	fmt.Println(validPassportCount)
}
