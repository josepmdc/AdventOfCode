package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var bagsContents map[string][]bag

type bag struct {
	quantity int
	color    string
}

func main() {
	data := util.GetData("7th/input")
	bagsContents = parseInput(data)
	part1()
	part2()
}

func part1() {
	count := 0
	for _, neighbours := range bagsContents {
		if contains(neighbours, "shiny gold") {
			count++
		}
	}
	fmt.Printf("Part 1: %d\n", count)
}

func part2() {
	shinyGold := bagsContents["shiny gold"]
	c := count(shinyGold)
	fmt.Printf("Part 2: %d\n", c)
}

func contains(neighbours []bag, color string) bool {
	for _, bag := range neighbours {
		if bag.color == color {
			return true
		}
		if contains(bagsContents[bag.color], color) {
			return true
		}
	}
	return false
}

func count(innerBags []bag) int {
	c := 0
	for _, bag := range innerBags {
		c += bag.quantity + bag.quantity*count(bagsContents[bag.color])
	}
	return c
}

func parseInput(data []string) map[string][]bag {
	bags := make(map[string][]bag)
	for _, rule := range data {
		colors := strings.Split(strings.Split(rule, "contain")[1], ",")
		contents := make([]bag, 0)
		for _, c := range colors {
			if !strings.Contains(c, "no other bags") {
				r := regexp.MustCompile(` (.*?) bag`)
				b := r.FindStringSubmatch(c)[1]
				bProperties := strings.SplitN(b, " ", 2)
				quantity, _ := strconv.Atoi(bProperties[0])
				contents = append(contents, bag{quantity: quantity, color: bProperties[1]})
			}
		}
		r := regexp.MustCompile(`(.*?) bag`)
		bags[r.FindStringSubmatch(rule)[1]] = contents
	}
	return bags
}
