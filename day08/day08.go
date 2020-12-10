package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation string

const (
	Acc = "acc"
	Jmp = "jmp"
	Nop = "nop"
)

type Instruction struct {
	op    Operation
	arg   int
	count int
}

// Given a rule string, parse it and create an instruction
func parseIns(line string) Instruction {
	re := regexp.MustCompile(`^(.+)\s(.+)$`)
	fields := re.FindStringSubmatch(line)
	arg, _ := strconv.Atoi(fields[2])
	return Instruction{Operation(fields[1]), arg, 0}
}

func validateProg(instrs []Instruction) (bool, int) {
	acc := 0
	index := 0
	var instr *Instruction
	for {
		instr = &instrs[index]
		//fmt.Printf("%+v: %d\n", instr, acc)
		if instr.count == 1 {
			return false, acc
		}
		switch instr.op {
		case Acc:
			acc += instr.arg
			index++
			instr.count += 1
			break
		case Jmp:
			index += instr.arg
			instr.count += 1
			break
		case Nop:
			index++
			instr.count += 1
			break
		}
		if index == len(instrs) {
			fmt.Println("Program ran successfully!")
			return true, acc
		}
	}
}

func main() {
	filename := os.Args[1]

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	s := strings.Split(string(input), "\n")

	var instrs []Instruction
	for _, line := range s {
		instr := parseIns(line)
		instrs = append(instrs, instr)
	}

	// Attempt to change each line and see if the program will run successfully
	for i, _ := range instrs {
		fmt.Printf("Switching operation %d\n", i)

		// Copy the instructions so we can modify the copy
		cpy := make([]Instruction, len(instrs))
		copy(cpy, instrs)

		if cpy[i].op == Jmp {
			cpy[i].op = Nop
		} else if cpy[i].op == Nop {
			cpy[i].op = Jmp
		} else {
			continue // Skip acc operations
		}

		isValid, acc := validateProg(cpy)
		if isValid {
			fmt.Printf("Success! Accumulator value is %d\n", acc)
			return
		}
	}
}
