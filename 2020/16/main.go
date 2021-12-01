package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//data := util.GetData("16/input_test")
	data := util.GetData("16/input")
	//part1(data[8:], data[0:3]) // for test input
	part1(data[25:], data[0:19])
}

func part1(tickets []string, rules []string) {
	sum := 0
	for _, ticket := range tickets {
		fields := strings.Split(ticket, ",")
		for _, field := range fields {
			value, _ := strconv.Atoi(field)
			valid := false
			for i := 0; i < len(rules) && !valid; i++ {
				fmt.Println(rules[i])
				fmt.Println(value)
				firstLower, firstUpper, secondLower, secondUpper := parseRule(rules[i])
				if (value >= firstLower && value <= firstUpper) || (value >= secondLower && value <= secondUpper) {
					valid = true
				}
			}
			if !valid {
				fmt.Println(value)
				sum += value
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func parseRule(rule string) (int, int, int, int) {
	parts := strings.Split(rule, " or ")
	r := regexp.MustCompile(`([0-9]+)-([0-9]+)`)
	first := r.FindStringSubmatch(parts[0])
	second := r.FindStringSubmatch(parts[1])
	return util.MustAtoi(first[1]), util.MustAtoi(first[2]),
		util.MustAtoi(second[1]), util.MustAtoi(second[2])
}
