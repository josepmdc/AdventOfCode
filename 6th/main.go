package main

import (
	"aoc2020/util"
	"fmt"
	"strings"
)

func main() {
	data := util.GetDataByBlocks("6th/input")
	part1(data)
	part2(data)
}

func part1(data []string) {
	result := 0

	for _, group := range data {
		answers := strings.Split(group, "\n")
		uniqueAnswers := getUniqueAnswers(answers)
		result += len(uniqueAnswers)
	}

	fmt.Printf("Part 1: %d\n", result)
}

func part2(data []string) {
	result := 0

	for _, group := range data {
		answers := strings.Split(group, "\n")
		uniqueAnswers := getUniqueAnswers(answers)

		for _, answer := range uniqueAnswers {
			if answer == len(answers) {
				result += 1
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func getUniqueAnswers(answers []string) map[rune]int {
	uniqueAnswers := make(map[rune]int)
	for _, answer := range answers {
		for _, letter := range answer {
			uniqueAnswers[letter] += 1
		}
	}
	return uniqueAnswers
}
