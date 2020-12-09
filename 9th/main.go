package main

import (
	"aoc2020/util"
	"fmt"
)

const preambleSize = 25

func main() {
	data := util.GetDataInt("9th/input")
	part1(data)
	part2(data)
}

func part1(data []int) {
	fmt.Printf("Part 1: %d\n", findInvalidNumber(data, preambleSize))
}

func part2(data []int) {
	num := findInvalidNumber(data, preambleSize)
	lower, upper := findSet(num, data)
	min, max := minMax(data[lower:upper])
	fmt.Printf("Part 2: min: %d, max: %d, min + max = %d\n", min, max, min+max)
}

func findInvalidNumber(data []int, preambleSize int) int {
	lower := 0
	for upper := preambleSize; upper < len(data); upper++ {
		preamble := data[lower:upper]
		valid := checkCombinations(data[upper], preamble)
		if !valid {
			return data[upper]
		}
		lower++
	}
	return 0
}

func checkCombinations(num int, data []int) bool {
	for _, num1 := range data {
		for _, num2 := range data {
			if num == num1+num2 {
				return true
			}
		}
	}
	return false
}

func findSet(num int, data []int) (int, int) {
	for i, lower := range data {
		sum := lower
		for j, upper := range data[i+1:] {
			if sum < num {
				sum += upper
			}
			if sum == num {
				return i, i + (j + 1)
			}
		}
	}
	return 0, 0
}

func minMax(values []int) (int, int) {
	min, max := values[0], values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
