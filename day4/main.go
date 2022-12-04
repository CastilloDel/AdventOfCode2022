package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
}

func part1() int {
	file, err := os.Open("day4/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		elf1, elf2 := parseElfAssignmentPair(scanner.Text())
		if elf1.contains(elf2) || elf2.contains(elf1) {
			total++
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

type ElfAssignment struct {
	start int
	end   int
}

func (assignment ElfAssignment) contains(otherAssignment ElfAssignment) bool {
	return assignment.start <= otherAssignment.start && assignment.end >= otherAssignment.end
}

func parseElfAssignmentPair(s string) (ElfAssignment, ElfAssignment) {
	elfAssignments := strings.Split(s, ",")
	return parseElfAssignment(elfAssignments[0]), parseElfAssignment(elfAssignments[1])
}

func parseElfAssignment(s string) ElfAssignment {
	limits := strings.Split(s, "-")
	start, err1 := strconv.Atoi(limits[0])
	end, err2 := strconv.Atoi(limits[1])
	if err1 != nil || err2 != nil {
		panic("Couldn't parse an assignment")
	}
	return ElfAssignment{start, end}
}
