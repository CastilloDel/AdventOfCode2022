package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
}

func part1() int {
	content, err := ioutil.ReadFile("day10/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	instructions := parseInstructions(string(content))
	state := ProgramState{1, 1}
	stop := 20
	finalStop := 220
	stopDistance := 40
	signalStrengthSum := 0
	for _, instruction := range instructions {
		newState := instruction.apply(state)
		if newState.programCounter > stop {
			signalStrengthSum += state.registerContent * stop
			if stop == finalStop {
				break
			}
			stop += stopDistance
		}
		state = newState
	}

	return signalStrengthSum
}

type ProgramState struct {
	programCounter  int
	registerContent int
}

type Instruction interface {
	apply(state ProgramState) ProgramState
}

type Addition struct {
	value int
}

func (addition Addition) apply(state ProgramState) ProgramState {
	state.programCounter += 2
	state.registerContent += addition.value
	return state
}

type NOOP struct{}

func (NOOP) apply(state ProgramState) ProgramState {
	state.programCounter++
	return state
}

func parseInstructions(s string) []Instruction {
	instructions := []Instruction{}
	for _, line := range strings.Split(s, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

func parseInstruction(line string) Instruction {
	if line == "noop" {
		return NOOP{}
	}
	value, err := strconv.Atoi(line[5:])
	if err != nil {
		panic("Couldn't read an instruction")
	}
	return Addition{value}
}
