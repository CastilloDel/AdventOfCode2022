package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
	fmt.Println("The result for the second part is:", part2())
}

func part1() int {
	content, err := ioutil.ReadFile("day9/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	instructions := parseInstructions(string(content))

	return countTailPositions(instructions, 2)
}

func part2() int {
	content, err := ioutil.ReadFile("day9/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	instructions := parseInstructions(string(content))

	return countTailPositions(instructions, 10)
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Instruction struct {
	direction Direction
	distance  int
}

func parseInstructions(s string) []Instruction {
	instructions := []Instruction{}
	for _, line := range strings.Split(s, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

func parseInstruction(s string) Instruction {
	var direction Direction
	switch s[0] {
	case 'U':
		direction = Up
	case 'R':
		direction = Right
	case 'D':
		direction = Down
	case 'L':
		direction = Left
	}
	distance, err := strconv.Atoi(s[2:])
	if err != nil {
		panic("Couldn't parse an Instruction")
	}
	return Instruction{direction: direction, distance: distance}
}

type Position struct {
	x int
	y int
}

func countTailPositions(instructions []Instruction, numberOfKnots int) int {
	totalPositions := map[Position]bool{}
	positions := []Position{}
	for i := 0; i < numberOfKnots; i++ {
		positions = append(positions, Position{0, 0})
	}
	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			positions[0] = moveKnot(positions[0], instruction.direction)
			for i := 1; i < len(positions); i++ {
				positions[i] = moveTail(positions[i-1], positions[i])
			}
			totalPositions[positions[len(positions)-1]] = true
		}
	}
	return len(totalPositions)
}

func moveKnot(position Position, direction Direction) Position {
	switch direction {
	case Up:
		position.y += 1
	case Right:
		position.x += 1
	case Down:
		position.y -= 1
	case Left:
		position.x -= 1
	}
	return position
}

func moveTail(head Position, tail Position) Position {
	distance := getManhattanDistance(head, tail)
	if distance > 2 { // we need to move in both directions
		if head.x-tail.x > 0 {
			tail.x += 1
		} else {
			tail.x -= 1
		}
		if head.y-tail.y > 0 {
			tail.y += 1
		} else {
			tail.y -= 1
		}
		// If they touch diagonally we don't need to move them
	} else if distance == 2 && int(math.Abs(float64(head.y-tail.y))) != 1 {
		xDistance := head.x - tail.x
		if xDistance == 2 {
			tail.x += 1
		} else if xDistance == -2 {
			tail.x -= 1
		}
		yDistance := head.y - tail.y
		if yDistance == 2 {
			tail.y += 1
		} else if yDistance == -2 {
			tail.y -= 1
		}
	}
	return tail
}

func getManhattanDistance(position1, position2 Position) int {
	return int(math.Abs(float64(position1.x-position2.x)) +
		math.Abs(float64(position1.y-position2.y)))
}
