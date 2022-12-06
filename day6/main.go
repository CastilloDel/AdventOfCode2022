package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
}

func part1() int {
	file, err := ioutil.ReadFile("day6/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	content := string(file)

	marker := []byte{content[0], content[1], content[2], content[3]}
	for i := 4; i < len(content); i++ {
		if checkAllDifferent(marker) {
			return i
		}
		marker = append(marker[1:], content[i])
	}
	panic("Couldn't find a marker")
}

func checkAllDifferent(marker []byte) bool {
	return marker[0] != marker[1] &&
		marker[0] != marker[2] &&
		marker[0] != marker[3] &&
		marker[1] != marker[2] &&
		marker[1] != marker[3] &&
		marker[2] != marker[3]
}
