package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
}

func part1() int {
	content, err := ioutil.ReadFile("day11/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	monkeys := parseInput(string(content))
	inspectionsPerMonkey := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		computeRound(monkeys, inspectionsPerMonkey)
	}
	max1 := 0
	max2 := 0
	for _, value := range inspectionsPerMonkey {
		if value > max1 {
			max2 = max1
			max1 = value
		} else if value > max2 {
			max2 = value
		}
	}
	return max1 * max2
}

func computeRound(monkeys []Monkey, inspections []int) []int {
	for monkeyIndex, monkey := range monkeys {
		for _, item := range monkey.items {
			item = Item(monkey.operation(int(item)) / 3)
			index := monkey.monkeyIfDivisible
			if int(item)%monkey.divisor != 0 {
				index = monkey.monkeyIfNotDivisible
			}
			monkeys[index].items = append(monkeys[index].items, item)
			inspections[monkeyIndex]++
		}
		// We can't use monkey directly, because it is a copy
		monkeys[monkeyIndex].items = []Item{}
	}
	return inspections
}

type Item int

type Monkey struct {
	operation            func(int) int
	divisor              int
	monkeyIfDivisible    int
	monkeyIfNotDivisible int
	items                []Item
}

func parseInput(s string) []Monkey {
	monkeys := []Monkey{}
	for _, monkey := range strings.Split(s, "\n\n") {
		monkeys = append(monkeys, parseMonkey(monkey))
	}
	return monkeys

}

// Monkey format:
// Monkey 1:
//
//	Starting items: 50, 99, 80, 84, 65, 95
//	Operation: new = old + 2
//	Test: divisible by 3
//	  If true: throw to monkey 4
//	  If false: throw to monkey 5
func parseMonkey(monkey string) Monkey {
	return Monkey{
		parseOperation(monkey),
		parseDivisor(monkey),
		parseMonkeyIfDivisible(monkey),
		parseMonkeyIfNotDivisible(monkey),
		parseItems(monkey),
	}
}

func parseOperation(monkey string) func(int) int {
	regex := regexp.MustCompile("Operation: new = (.*)")
	operationString := regex.FindStringSubmatch(monkey)[1]
	operationParts := strings.Split(operationString, " ")
	return func(value int) int {
		operand1, err := strconv.Atoi(operationParts[0])
		if err != nil {
			operand1 = value
		}
		operand2, err := strconv.Atoi(operationParts[2])
		if err != nil {
			operand2 = value
		}
		if operationParts[1] == "+" {
			return operand1 + operand2
		} else {
			return operand1 * operand2
		}
	}
}

func parseDivisor(monkey string) int {
	regex := regexp.MustCompile("Test: divisible by (\\d+)")
	divisorString := regex.FindStringSubmatch(monkey)[1]
	divisor, _ := strconv.Atoi(divisorString)
	return divisor
}

func parseMonkeyIfDivisible(monkey string) int {
	regex := regexp.MustCompile("If true: throw to monkey (\\d+)")
	monkeyIndexString := regex.FindStringSubmatch(monkey)[1]
	monkeyIndex, _ := strconv.Atoi(monkeyIndexString)
	return monkeyIndex
}

func parseMonkeyIfNotDivisible(monkey string) int {
	regex := regexp.MustCompile("If false: throw to monkey (\\d+)")
	monkeyIndexString := regex.FindStringSubmatch(monkey)[1]
	monkeyIndex, _ := strconv.Atoi(monkeyIndexString)
	return monkeyIndex
}

func parseItems(monkey string) []Item {
	regex := regexp.MustCompile("Starting items: ((?:\\d+(?:, )?)*)")
	itemsString := regex.FindStringSubmatch(monkey)[1]
	items := []Item{}
	for _, item := range strings.Split(itemsString, ", ") {
		result, err := strconv.Atoi(item)
		if err != nil {
			panic("Couldn't parse")
		}
		items = append(items, Item(result))
	}
	return items
}
