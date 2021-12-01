package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := util.GetData("2nd/input")
	part1(data)
	part2(data)
}

func part1(data []string) {
	validPasswords := 0

	for _, line := range data {
		password, letter, lower, upper := parseLine(line)

		count := strings.Count(password, letter)
		if count >= lower && count <= upper {
			validPasswords++
		}
	}

	fmt.Printf("Part 1: There are %d valid passwords\n", validPasswords)
}

func part2(data []string) {
	validPasswords := 0

	for _, line := range data {
		password, letter, lower, upper := parseLine(line)

		if (string(password[lower]) == letter) != (string(password[upper]) == letter) {
			validPasswords++
		}
	}

	fmt.Printf("Part 2: There are %d valid passwords\n", validPasswords)
}

func parseLine(line string) (string, string, int, int) {
	splitLine := strings.Split(line, ":")
	passwordPolicy := strings.Split(splitLine[0], " ")

	password := splitLine[1]
	letter := passwordPolicy[1]

	validRange := strings.Split(passwordPolicy[0], "-")

	lower, err := strconv.Atoi(validRange[0])
	if err != nil {
		fmt.Println(err)
	}

	upper, err := strconv.Atoi(validRange[1])
	if err != nil {
		fmt.Println(err)
	}

	return password, letter, lower, upper
}
