package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
	fmt.Println("The result for the second part is:", part2())
}

func part1() int {
	content, err := ioutil.ReadFile("day6/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	marker := content[0:4]
	for i := 4; i < len(content); i++ {
		if checkAllDifferent(marker) {
			return i
		}
		marker = append(marker[1:], content[i])
	}
	panic("Couldn't find an start marker")
}

func part2() int {
	content, err := ioutil.ReadFile("day6/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	marker := content[0:14]
	for i := 14; i < len(content); i++ {
		if checkAllDifferent(marker) {
			return i
		}
		marker = append(marker[1:], content[i])
	}
	panic("Couldn't find a message marker")
}

func checkAllDifferent(marker []byte) bool {
	appearances := make(map[byte]bool)
	for _, value := range marker {
		_, ok := appearances[value]
		if ok {
			return false
		}
		appearances[value] = true
	}
	return true
}
