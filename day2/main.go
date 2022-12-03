package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
	fmt.Println("The result for the second part is:", part2())
}

func part1() int {
	file, err := os.Open("day2/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		plays := strings.Fields(scanner.Text())
		rival_play := readPlay(plays[0])
		own_play := readPlay(plays[1])
		total += int(own_play)
		total += int(getMatchResult(own_play, rival_play))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func part2() int {
	file, err := os.Open("day2/input")
	if err != nil {
		panic("Couldn't read the input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		rival_play := readPlay(words[0])
		result := readResult(words[1])
		total += int(result)
		total += int(getPlayFromResult(result, rival_play))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

type Play int64

const (
	Rock     Play = 1
	Paper         = 2
	Scissors      = 3
)

func readPlay(s string) Play {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	panic("Couldn't read a play")
}

type Result int64

const (
	Win  Result = 6
	Draw        = 3
	Lose        = 0
)

func readResult(s string) Result {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	panic("Couldn't read a result")
}

func getMatchResult(play Play, rival_play Play) Result {
	if play == rival_play {
		return Win
	}
	if play == Paper && rival_play == Rock ||
		play == Scissors && rival_play == Paper ||
		play == Rock && rival_play == Scissors {
		return Draw
	}
	return Lose
}

func getPlayFromResult(result Result, rival_play Play) Play {
	if result == Draw {
		return rival_play
	} else if result == Win {
		if rival_play == Rock {
			return Paper
		} else if rival_play == Paper {
			return Scissors
		}
		return Rock
	}
	if rival_play == Rock {
		return Scissors
	} else if rival_play == Paper {
		return Rock
	}
	return Paper
}
