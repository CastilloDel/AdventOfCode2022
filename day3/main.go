package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRucksackSharedItem(rucksack string) string {
	first_compartment := make(map[byte]int)
	for i := 0; i < len(rucksack)/2; i++ {
		first_compartment[rucksack[i]] = 1
	}
	for i := len(rucksack) / 2; i < len(rucksack); i++ {
		if first_compartment[rucksack[i]] == 1 {
			return string(rucksack[i])
		}
	}
	panic("A rucksack didn't have a shared item")
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func getItemPriority(item string) int {
	ascii_value := int(item[0])
	if ascii_value >= int('a') && ascii_value <= int('z') {
		return ascii_value - int('a') + 1
	}
	return ascii_value - int('A') + 27
}

func part1() int {
	file, err := os.Open("day3/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		sharedItem := getRucksackSharedItem(rucksack)
		total += getItemPriority(sharedItem)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func main() {
	fmt.Println("The result for the first part is:", part1())
}
