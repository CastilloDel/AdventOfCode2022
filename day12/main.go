package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type ProblemData struct {
	grid  Grid
	start Position
	end   Position
}

func main() {
	content, err := ioutil.ReadFile("day12/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	input := parseInput(string(content))

	fmt.Println("The result for the first part is:", part1(input))
	fmt.Println("The result for the second part is:", part2(input))
}

func parseInput(s string) ProblemData {
	grid := [][]int{}
	start := Position{}
	end := Position{}
	for y, row := range strings.Split(s, "\n") {
		grid = append(grid, []int{})
		for x, position := range strings.Split(row, "") {
			if position == "S" {
				start = Position{x, y}
				position = "a"
			} else if position == "E" {
				end = Position{x, y}
				position = "z"
			}
			grid[y] = append(grid[y], int(position[0])-int('a'))
		}
	}
	return ProblemData{grid, start, end}
}

func part1(input ProblemData) int {
	return findMinimumSteps(input.grid, input.start, input.end)
}

func part2(input ProblemData) int {
	min := math.MaxInt
	for _, start := range findPositionsWithHeight(input.grid, 0) {
		minimumSteps := findMinimumSteps(input.grid, start, input.end)
		if minimumSteps < min {
			min = minimumSteps
		}
	}
	return min
}

type Position struct {
	x int
	y int
}

type Grid = [][]int

type PossiblePath struct {
	current        Position
	stepsTaken     int
	stepsEstimated int
}

func findMinimumSteps(grid Grid, start, end Position) int {
	positions := map[Position]int{}
	minimums := map[Position]int{}
	positions[start] = manhattanDistance(grid, start, end)
	minimums[start] = 0
	for {
		position := findPromisingPath(positions)
		if position == end {
			return minimums[position]
		}
		for _, neighbor := range getNeighbors(grid, position) {
			min, present := minimums[neighbor]
			if present && min <= minimums[position]+1 {
				continue
			}
			minimums[neighbor] = minimums[position] + 1
			positions[neighbor] = minimums[neighbor] + manhattanDistance(grid, neighbor, end)
		}

	}
}

func findPromisingPath(positions map[Position]int) Position {
	best := Position{}
	bestValue := math.MaxInt
	for position, value := range positions {
		if value < bestValue {
			best = position
			bestValue = value
		}
	}
	delete(positions, best)
	return best
}

func getNeighbors(grid Grid, pos Position) []Position {
	possibleNeighbors := []Position{
		{pos.x - 1, pos.y}, {pos.x + 1, pos.y}, {pos.x, pos.y - 1}, {pos.x, pos.y + 1},
	}
	neighbors := []Position{}
	for _, possible := range possibleNeighbors {
		if possible.x < 0 || possible.y < 0 ||
			possible.x >= len(grid[0]) || possible.y >= len(grid) {
			continue
		}
		if grid[pos.y][pos.x]-grid[possible.y][possible.x] >= -1 {
			neighbors = append(neighbors, possible)
		}
	}
	return neighbors
}

func manhattanDistance(grid Grid, pos1, pos2 Position) int {
	return int(math.Abs(float64(pos1.x-pos2.x)) + math.Abs(float64(pos1.y-pos2.y)))
}

func findPositionsWithHeight(grid Grid, desiredHeight int) []Position {
	positions := []Position{}
	for y, row := range grid {
		for x, height := range row {
			if height == desiredHeight {
				positions = append(positions, Position{x, y})
			}

		}
	}
	return positions
}
