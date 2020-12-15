package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Printf("Part 1: %d\n", solve(2020))
}

func part2() {
	fmt.Printf("Part 2: %d\n", solve(30000000))
}

func solve(index int) int {
	cache := map[int]int{11: 1, 18: 2, 0: 3, 20: 4, 1: 5, 7: 6}
	last := 16
	n := 0
	for i := 8; i <= index; i++ {
		if cache[last] != 0 {
			n = i - 1 - cache[last]
		} else {
			n = 0
		}
		cache[last] = i - 1
		last = n
	}
	return n
}
