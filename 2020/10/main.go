package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

func main() {
	adapters := util.GetDataInt("10th/input")
	sort.Ints(adapters)
	part1(adapters)
	part2(adapters)
}

func part1(adapters []int) {
	diff1 := 1
	diff3 := 1
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		if diff == 1 {
			diff1++
		} else {
			diff3++
		}
	}
	fmt.Printf("Part 1: diff 1: %d, diff 3: %d, diff1 * diff3 = %d\n", diff1, diff3, diff1*diff3)
}

func part2(adapters []int) {
	joltages := map[int]int{0: 1}
	for _, adapter := range adapters {
		joltages[adapter] = joltages[adapter-1] + joltages[adapter-2] + joltages[adapter-3]
	}
	fmt.Printf("Part 2: %d\n", joltages[adapters[len(adapters)-1]])
}
