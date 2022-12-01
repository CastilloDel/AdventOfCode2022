package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	max := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if count > max {
				max = count
			}
			count = 0
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		count += number
	}

	fmt.Println("The result for the first part is:", max)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
