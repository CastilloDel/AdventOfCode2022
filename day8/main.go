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
	content, err := ioutil.ReadFile("day8/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	treeHeights := parseTreeHeights(string(content))
	treeVisibilities := getTreeVisibilities(treeHeights)

	totalVisibleTrees := 0
	for _, row := range treeVisibilities {
		for _, visibility := range row {
			if visibility {
				totalVisibleTrees++
			}
		}
	}

	return totalVisibleTrees
}

func parseTreeHeights(content string) [][]int {
	treeHeights := make([][]int, 0)
	for i, line := range strings.Split(content, "\n") {
		treeHeights = append(treeHeights, make([]int, 0))
		for _, unparsedHeight := range line {
			height, err := strconv.Atoi(string(unparsedHeight))
			if err != nil {
				panic("Couldn't parse a tree")
			}
			treeHeights[i] = append(treeHeights[i], height)
		}
	}
	return treeHeights
}

func getTreeVisibilities(treeHeights [][]int) [][]bool {
	m := len(treeHeights)
	n := len(treeHeights[0])
	treeVisibilities := make([][]bool, len(treeHeights))
	for i := 0; i < m; i++ { // Left -> Right
		treeVisibilities[i] = make([]bool, n)
		treeVisibilities[i][0] = true
		currentHeight := treeHeights[i][0]
		for j := 1; j < n && currentHeight < 9; j++ {
			if treeHeights[i][j] > currentHeight {
				currentHeight = treeHeights[i][j]
				treeVisibilities[i][j] = true
			}
		}
	}
	for i := 0; i < m; i++ { // Right -> Left
		treeVisibilities[i][n-1] = true
		currentHeight := treeHeights[i][n-1]
		for j := n - 2; j >= 0 && currentHeight < 9; j-- {
			if treeHeights[i][j] > currentHeight {
				currentHeight = treeHeights[i][j]
				treeVisibilities[i][j] = true
			}
		}
	}
	for j := 0; j < n; j++ { // Top -> Down
		treeVisibilities[0][j] = true
		currentHeight := treeHeights[0][j]
		for i := 1; i < m && currentHeight < 9; i++ {
			if treeHeights[i][j] > currentHeight {
				currentHeight = treeHeights[i][j]
				treeVisibilities[i][j] = true
			}
		}
	}
	for j := 0; j < n; j++ { // Bottom -> Up
		treeVisibilities[m-1][j] = true
		currentHeight := treeHeights[m-1][j]
		for i := m - 2; i >= 9 && currentHeight < 9; i-- {
			if treeHeights[i][j] > currentHeight {
				currentHeight = treeHeights[i][j]
				treeVisibilities[i][j] = true
			}
		}
	}

	return treeVisibilities
}
