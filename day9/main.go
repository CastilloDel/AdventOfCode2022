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
}

func part1() int {
	content, err := ioutil.ReadFile("day9/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	instructions := parseInstructions(string(content))

	return countTailPositions(instructions)
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

type RopePosition struct {
	head Position
	tail Position
}

func countTailPositions(instructions []Instruction) int {
	totalPositions := map[Position]bool{}
	position := RopePosition{Position{0, 0}, Position{0, 0}}
	for _, instruction := range instructions {
		for i := 0; i < instruction.distance; i++ {
			position = moveRope(position, instruction.direction)
			totalPositions[position.tail] = true
		}
	}
	return len(totalPositions)
}

func moveRope(position RopePosition, direction Direction) RopePosition {
	switch direction {
	case Up:
		position.head.y += 1
	case Right:
		position.head.x += 1
	case Down:
		position.head.y -= 1
	case Left:
		position.head.x -= 1
	}
	return moveTail(position)
}

func moveTail(position RopePosition) RopePosition {
	distance := getManhattanDistance(position.head, position.tail)
	if distance == 3 { // we need to move in both directions
		if position.head.x-position.tail.x > 0 {
			position.tail.x += 1
		} else {
			position.tail.x -= 1
		}
		if position.head.y-position.tail.y > 0 {
			position.tail.y += 1
		} else {
			position.tail.y -= 1
		}
		// If they touch diagonally we don't need to move them
	} else if distance == 2 && int(math.Abs(float64(position.head.y-position.tail.y))) != 1 {
		xDistance := position.head.x - position.tail.x
		if xDistance == 2 {
			position.tail.x += 1
		} else if xDistance == -2 {
			position.tail.x -= 1
		}
		yDistance := position.head.y - position.tail.y
		if yDistance == 2 {
			position.tail.y += 1
		} else if yDistance == -2 {
			position.tail.y -= 1
		}
	}
	return position
}

func getManhattanDistance(position1, position2 Position) int {
	return int(math.Abs(float64(position1.x-position2.x)) +
		math.Abs(float64(position1.y-position2.y)))
}
