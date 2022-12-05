package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
}

type Instruction struct {
	size   int
	source int
	dest   int
}

func part1() string {
	file, err := os.Open("day5/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	arrangement := parseArrangement(scanner)
	scanner.Scan()
	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		applyInstruction(arrangement, instruction)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := ""
	for i := 0; i < 9; i++ {
		result += arrangement[i][len(arrangement[i])-1]
	}

	return result
}

func parseArrangement(scanner *bufio.Scanner) map[int][]string {
	arrangement := make(map[int][]string)
	scanner.Scan()
	line := scanner.Text()
	for line[1] != byte('1') {
		for i := 0; i*4 < len(line); i++ {
			if line[i*4] == '[' && line[i*4+2] == ']' {
				_, ok := arrangement[i]
				if ok {
					arrangement[i] = append([]string{string(line[i*4+1])}, arrangement[i]...)
				} else {
					arrangement[i] = []string{string(line[i*4+1])}
				}
			}
		}
		scanner.Scan()
		line = scanner.Text()
	}
	return arrangement
}

func parseInstruction(s string) Instruction {
	regex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	unparsedParts := regex.FindStringSubmatch(s)
	var parts [3]int
	for i, value := range unparsedParts[1:] {
		result, err := strconv.Atoi(value)
		if err != nil {
			panic("Couldn't parse instruction")
		}
		parts[i] = result
	}
	// We substract 1, because they are 1 based indexes
	return Instruction{parts[0], parts[1] - 1, parts[2] - 1}
}

func applyInstruction(arrangement map[int][]string, instruction Instruction) {
	source := instruction.source
	dest := instruction.dest
	for i := 0; i < instruction.size; i++ {
		sourceLen := len(arrangement[source])
		lastElement := arrangement[source][sourceLen-1]
		arrangement[dest] = append(arrangement[dest], lastElement)
		arrangement[source] = arrangement[source][:sourceLen-1]
	}
}
