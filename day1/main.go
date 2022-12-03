package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
	fmt.Println("The result for the second part is:", part2())
}

func part1() int {
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return max
}

func part2() int {
	file, err := os.Open("day1/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxThree := [3]int{0, 0, 0}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if count > maxThree[0] {
				maxThree[0] = count
			}
			if maxThree[0] > maxThree[1] {
				maxThree[0], maxThree[1] = maxThree[1], maxThree[0]
			}
			if maxThree[1] > maxThree[2] {
				maxThree[1], maxThree[2] = maxThree[2], maxThree[1]
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return maxThree[0] + maxThree[1] + maxThree[2]
}
